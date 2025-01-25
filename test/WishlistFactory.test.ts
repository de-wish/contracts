import { ethers } from "hardhat";
import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { expect } from "chai";
import { Reverter } from "@test-helpers";
import { PRECISION, wei } from "@scripts";
import { WishlistFactory, Wishlist, ERC20Mock } from "@ethers-v6";

describe("WishlistFactory", () => {
  const reverter = new Reverter();

  const initialTokensAmount = wei(10000);
  const defaultProtocolFeePercentage = 2n * PRECISION;
  const maxProtocolFeeAmount = wei(100, 6);

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let FEE_RECIPIENT: SignerWithAddress;

  let usdcToken: ERC20Mock;

  let wishlistImpl: Wishlist;
  let wishlistFactoryImpl: WishlistFactory;

  let wishlistFactory: WishlistFactory;

  before(async () => {
    [OWNER, FIRST, SECOND, FEE_RECIPIENT] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory("ERC1967Proxy");
    const WishlistFactoryFactory = await ethers.getContractFactory("WishlistFactory");
    const WishlistFactory = await ethers.getContractFactory("Wishlist");
    const ERC20MockFactory = await ethers.getContractFactory("ERC20Mock");

    usdcToken = await ERC20MockFactory.deploy("Mock USDC", "USDC", 6);

    await usdcToken.mint(OWNER, initialTokensAmount);
    await usdcToken.mint(FIRST, initialTokensAmount);

    wishlistFactoryImpl = await WishlistFactoryFactory.deploy();
    wishlistImpl = await WishlistFactory.deploy();

    const wishlistFactoryProxy = await ERC1967ProxyFactory.deploy(wishlistFactoryImpl, "0x");

    wishlistFactory = await ethers.getContractAt("WishlistFactory", wishlistFactoryProxy);

    await wishlistFactory.initialize(wishlistImpl, usdcToken, defaultProtocolFeePercentage, maxProtocolFeeAmount);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("initialize", () => {
    it("should set parameters correctly", async () => {
      expect(await wishlistFactory.owner()).to.eq(OWNER);
      expect(await wishlistFactory.usdcToken()).to.eq(usdcToken);
      expect(await wishlistFactory.protocolFeePercentage()).to.eq(defaultProtocolFeePercentage);
      expect(await wishlistFactory.maxProtocolFeeAmount()).to.eq(maxProtocolFeeAmount);

      const proxyBeacon = await ethers.getContractAt("ProxyBeacon", await wishlistFactory.proxyBeacon());

      expect(await proxyBeacon.implementation()).to.be.eq(wishlistImpl);
    });

    it("should get exception if try to call init function twice", async () => {
      await expect(
        wishlistFactory.initialize(wishlistImpl, usdcToken, defaultProtocolFeePercentage, maxProtocolFeeAmount),
      ).to.be.revertedWith("Initializable: contract is already initialized");
    });
  });

  describe("upgradability", () => {
    it("should correctly upgrade WishlistFactory contract", async () => {
      const WishlistFactoryMockFactory = await ethers.getContractFactory("WishlistFactoryMock");

      const wishlistFactoryMockImpl = await WishlistFactoryMockFactory.deploy();

      const newWishlistFactory = await ethers.getContractAt("WishlistFactoryMock", wishlistFactory);

      await expect(newWishlistFactory.version()).to.be.revertedWithoutReason();

      await wishlistFactory.upgradeTo(wishlistFactoryMockImpl);

      expect(await newWishlistFactory.version()).to.be.eq("1.1.1");
    });

    it("should get exception if non-owner try to upgrade WishlistFactory contract", async () => {
      await expect(wishlistFactory.connect(FIRST).upgradeTo(ethers.ZeroAddress)).to.revertedWith(
        "Ownable: caller is not the owner",
      );
    });
  });

  describe("upgradeWishlistsImpl", () => {
    it("should correctly upgrade wishlists implementation", async () => {
      const wishlistAddr = await wishlistFactory.createWishlist.staticCall([]);

      await wishlistFactory.createWishlist([]);

      const WishlistMockFactory = await ethers.getContractFactory("WishlistMock");

      const wishlistMockImpl = await WishlistMockFactory.deploy();

      const newWishlist = await ethers.getContractAt("WishlistMock", wishlistAddr);

      await expect(newWishlist.version()).to.be.revertedWithoutReason();

      await wishlistFactory.upgradeWishlistsImpl(wishlistMockImpl);

      const proxyBeacon = await ethers.getContractAt("ProxyBeacon", await wishlistFactory.proxyBeacon());

      expect(await proxyBeacon.implementation()).to.be.eq(wishlistMockImpl);

      expect(await newWishlist.version()).to.be.eq("2.0.0");
    });

    it("should get exception if non-owner try to upgrade wishlists implementation", async () => {
      await expect(wishlistFactory.connect(FIRST).upgradeWishlistsImpl(ethers.ZeroAddress)).to.revertedWith(
        "Ownable: caller is not the owner",
      );
    });
  });

  describe("setProtocolFeePercentage", () => {
    it("should correctly set new protocol fee value", async () => {
      const newProtocolFee = 3n * PRECISION;

      const tx = await wishlistFactory.setProtocolFeePercentage(newProtocolFee);

      expect(tx).to.emit(wishlistFactory, "ProtocolFeeUpdated").withArgs(newProtocolFee, defaultProtocolFeePercentage);

      expect(await wishlistFactory.protocolFeePercentage()).to.be.eq(newProtocolFee);
    });

    it("should get exception if try to set protocol fee higher than the max value", async () => {
      const newProtocolFee = 110n * PRECISION;

      await expect(wishlistFactory.setProtocolFeePercentage(newProtocolFee))
        .to.revertedWithCustomError(wishlistFactory, "InvalidProtocolFee")
        .withArgs(newProtocolFee);
    });

    it("should get exception if non-owner try to set protocol fee", async () => {
      const newProtocolFee = 3n * PRECISION;

      await expect(wishlistFactory.connect(FIRST).setProtocolFeePercentage(newProtocolFee)).to.revertedWith(
        "Ownable: caller is not the owner",
      );
    });
  });

  describe("withdrawProtocolFee", () => {
    const feeAmount = wei(100, 6);

    beforeEach("setup", async () => {
      await usdcToken.transfer(wishlistFactory, feeAmount);
    });

    it("should correctly withdraw protocol fee from the contract", async () => {
      const tx = await wishlistFactory.withdrawProtocolFee(FEE_RECIPIENT);

      expect(tx).to.emit(wishlistFactory, "ProtocolFeeWithdrawn").withArgs(feeAmount, FEE_RECIPIENT);
      expect(await usdcToken.balanceOf(FEE_RECIPIENT)).to.be.eq(feeAmount);

      expect(await usdcToken.balanceOf(wishlistFactory)).to.be.eq(0);
    });

    it("should get exception if pass zero recipient address", async () => {
      await expect(wishlistFactory.withdrawProtocolFee(ethers.ZeroAddress))
        .to.revertedWithCustomError(wishlistFactory, "ZeroFeeRecipient")
        .withArgs();
    });

    it("should get exception if nothing to withdraw", async () => {
      await wishlistFactory.withdrawProtocolFee(FEE_RECIPIENT);

      await expect(wishlistFactory.withdrawProtocolFee(FEE_RECIPIENT))
        .to.revertedWithCustomError(wishlistFactory, "ZeroFeeToWithdraw")
        .withArgs();
    });

    it("should get exception if non-owner try to withdraw protocol fee", async () => {
      await expect(wishlistFactory.connect(FIRST).withdrawProtocolFee(FIRST)).to.revertedWith(
        "Ownable: caller is not the owner",
      );
    });
  });
});
