import { ethers } from "hardhat";
import { BigNumberish } from "ethers";
import { expect } from "chai";

import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { time } from "@nomicfoundation/hardhat-network-helpers";

import { PRECISION, wei } from "@scripts";
import { getCreateWishlistSignature, getItemPricesHash, Reverter } from "@test-helpers";

import { WishlistFactory, Wishlist, ERC20Mock, IWishlist, ProtocolTreasury } from "@ethers-v6";

describe("WishlistFactory", () => {
  const reverter = new Reverter();

  const initialTokensAmount = wei(10000);
  const defaultProtocolFeePercentage = 2n * PRECISION;
  const maxProtocolFeeAmount = wei(100, 6);

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let PROTOCOL_SIGNER: SignerWithAddress;

  let usdcToken: ERC20Mock;

  let wishlistImpl: Wishlist;
  let wishlistFactoryImpl: WishlistFactory;

  let protocolTreasury: ProtocolTreasury;
  let wishlistFactory: WishlistFactory;

  function checkItemInfo(
    itemInfo: IWishlist.WishlistItemInfoStruct,
    expectedItemId: BigNumberish,
    expectedPrice: BigNumberish,
    expectedCollectedAmount: BigNumberish = 0,
    expectedContributionsInfo: IWishlist.ContributionInfoStruct[] = [],
  ) {
    expect(itemInfo.itemId).to.be.eq(expectedItemId);
    expect(itemInfo.itemPrice).to.be.eq(expectedPrice);
    expect(itemInfo.collectedTokensAmount).to.be.eq(expectedCollectedAmount);

    expect(itemInfo.totalContributionsNumber).to.be.eq(expectedContributionsInfo.length);

    for (let i = 0; i < expectedContributionsInfo.length; i++) {
      expect(itemInfo.contributionsInfo[i].contributionId).to.be.eq(expectedContributionsInfo[i].contributionId);
      expect(itemInfo.contributionsInfo[i].contributor).to.be.eq(expectedContributionsInfo[i].contributor);
      expect(itemInfo.contributionsInfo[i].contributionId).to.be.eq(expectedContributionsInfo[i].contributionAmount);
    }

    expect(itemInfo.isActive).to.be.eq(expectedCollectedAmount !== expectedPrice);
  }

  before(async () => {
    [OWNER, FIRST, SECOND, PROTOCOL_SIGNER] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory("ERC1967Proxy");
    const ProtocolTreasuryFactory = await ethers.getContractFactory("ProtocolTreasury");
    const WishlistFactoryFactory = await ethers.getContractFactory("WishlistFactory");
    const WishlistFactory = await ethers.getContractFactory("Wishlist");
    const ERC20MockFactory = await ethers.getContractFactory("ERC20Mock");

    usdcToken = await ERC20MockFactory.deploy("Mock USDC", "USDC", 6);

    await usdcToken.mint(OWNER, initialTokensAmount);
    await usdcToken.mint(FIRST, initialTokensAmount);

    const protocolTreasuryImpl = await ProtocolTreasuryFactory.deploy();

    wishlistFactoryImpl = await WishlistFactoryFactory.deploy();
    wishlistImpl = await WishlistFactory.deploy();

    const protocolTreasuryProxy = await ERC1967ProxyFactory.deploy(protocolTreasuryImpl, "0x");
    const wishlistFactoryProxy = await ERC1967ProxyFactory.deploy(wishlistFactoryImpl, "0x");

    protocolTreasury = await ethers.getContractAt("ProtocolTreasury", protocolTreasuryProxy);
    wishlistFactory = await ethers.getContractAt("WishlistFactory", wishlistFactoryProxy);

    await protocolTreasury.initialize();
    await wishlistFactory.initialize(wishlistImpl, usdcToken, PROTOCOL_SIGNER, {
      protocolFeePercentage: defaultProtocolFeePercentage,
      maxProtocolFeeAmount: maxProtocolFeeAmount,
      protocolFeeRecipient: protocolTreasury,
    });

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("initialize", () => {
    it("should set parameters correctly", async () => {
      expect(await wishlistFactory.owner()).to.eq(OWNER);
      expect(await wishlistFactory.usdcToken()).to.eq(usdcToken);
      expect(await wishlistFactory.getProtocolFeeSettings()).to.deep.eq([
        defaultProtocolFeePercentage,
        maxProtocolFeeAmount,
        await protocolTreasury.getAddress(),
      ]);

      const proxyBeacon = await ethers.getContractAt("ProxyBeacon", await wishlistFactory.proxyBeacon());

      expect(await proxyBeacon.implementation()).to.be.eq(wishlistImpl);
    });

    it("should get exception if try to call init function twice", async () => {
      await expect(
        wishlistFactory.initialize(wishlistImpl, usdcToken, PROTOCOL_SIGNER, {
          protocolFeePercentage: defaultProtocolFeePercentage,
          maxProtocolFeeAmount: maxProtocolFeeAmount,
          protocolFeeRecipient: protocolTreasury,
        }),
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
      const wishlistId = 123;
      const sigDeadline = (await time.latest()) + 1000;

      const signature = await getCreateWishlistSignature(
        PROTOCOL_SIGNER,
        await FIRST.getAddress(),
        wishlistId,
        getItemPricesHash([]),
        sigDeadline,
        await wishlistFactory.getAddress(),
      );

      await wishlistFactory.connect(PROTOCOL_SIGNER).createWishlist(FIRST, wishlistId, sigDeadline, [], signature);

      const newWishlistAddr = await wishlistFactory.getWishlistAddress(wishlistId);

      const WishlistMockFactory = await ethers.getContractFactory("WishlistMock");

      const wishlistMockImpl = await WishlistMockFactory.deploy();

      const newWishlist = await ethers.getContractAt("WishlistMock", newWishlistAddr);

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

  describe("setProtocolSignerAddr", () => {
    it("should correctly set protocol signer address", async () => {
      const tx = await wishlistFactory.setProtocolSignerAddr(FIRST);

      expect(tx).to.emit(wishlistFactory, "ProtocolSignerUpdated").withArgs(FIRST, PROTOCOL_SIGNER);

      expect(await wishlistFactory.protocolSignerAddr()).to.be.eq(FIRST);
    });

    it("should get exception if try to set zero address as a protocol signer address", async () => {
      await expect(wishlistFactory.setProtocolSignerAddr(ethers.ZeroAddress))
        .to.revertedWithCustomError(wishlistFactory, "ZeroAddr")
        .withArgs("ProtocolSigner");
    });

    it("should get exception if non-owner try to set protocol signer address", async () => {
      await expect(wishlistFactory.connect(FIRST).setProtocolSignerAddr(FIRST)).to.revertedWith(
        "Ownable: caller is not the owner",
      );
    });
  });

  describe("setProtocolFeeSettings", () => {
    it("should correctly set protocol fee settings", async () => {
      const newProtocolFee = 3n * PRECISION;

      const tx = await wishlistFactory.setProtocolFeeSettings({
        protocolFeePercentage: newProtocolFee,
        maxProtocolFeeAmount: maxProtocolFeeAmount,
        protocolFeeRecipient: FIRST,
      });

      expect(tx)
        .to.emit(wishlistFactory, "ProtocolFeeSettingsUpdated")
        .withArgs(newProtocolFee, maxProtocolFeeAmount, FIRST);
    });

    it("should get exception if try to set protocol fee higher than the max value", async () => {
      const newProtocolFee = 110n * PRECISION;

      await expect(
        wishlistFactory.setProtocolFeeSettings({
          protocolFeePercentage: newProtocolFee,
          maxProtocolFeeAmount: maxProtocolFeeAmount,
          protocolFeeRecipient: FIRST,
        }),
      )
        .to.revertedWithCustomError(wishlistFactory, "InvalidProtocolFee")
        .withArgs(newProtocolFee);
    });

    it("should get exception if try to set zero address as a protocol fee recipient", async () => {
      const newProtocolFee = 3n * PRECISION;

      await expect(
        wishlistFactory.setProtocolFeeSettings({
          protocolFeePercentage: newProtocolFee,
          maxProtocolFeeAmount: maxProtocolFeeAmount,
          protocolFeeRecipient: ethers.ZeroAddress,
        }),
      )
        .to.revertedWithCustomError(wishlistFactory, "ZeroAddr")
        .withArgs("ProtocolFeeRecipient");
    });

    it("should get exception if non-owner try to set protocol fee settings", async () => {
      const newProtocolFee = 3n * PRECISION;

      await expect(
        wishlistFactory.connect(FIRST).setProtocolFeeSettings({
          protocolFeePercentage: newProtocolFee,
          maxProtocolFeeAmount: maxProtocolFeeAmount,
          protocolFeeRecipient: FIRST,
        }),
      ).to.revertedWith("Ownable: caller is not the owner");
    });
  });

  describe("createWishlist", () => {
    it("should correctly create wishlist", async () => {
      const wishlistId = 123;
      const itemPrices = [wei(100, 6), wei(120, 6)];

      const sigDeadline = (await time.latest()) + 1000;

      const signature = await getCreateWishlistSignature(
        PROTOCOL_SIGNER,
        await FIRST.getAddress(),
        wishlistId,
        getItemPricesHash(itemPrices),
        sigDeadline,
        await wishlistFactory.getAddress(),
      );

      const tx = await wishlistFactory
        .connect(PROTOCOL_SIGNER)
        .createWishlist(FIRST, wishlistId, sigDeadline, itemPrices, signature);

      const newWishlistAddr = await wishlistFactory.getWishlistAddress(wishlistId);

      await expect(tx).to.emit(wishlistFactory, "NewWishlistCreated").withArgs(FIRST, wishlistId, newWishlistAddr);

      const newWishlist = await ethers.getContractAt("Wishlist", newWishlistAddr);

      expect(await newWishlist.wishlistId()).to.be.eq(wishlistId);
      expect(await newWishlist.wishlistOwner()).to.be.eq(FIRST);
      expect(await wishlistFactory.wishlistExists(wishlistId)).to.be.true;
      expect(await wishlistFactory.getWishlistsTotalCount()).to.be.eq(1);

      const info = await newWishlist.getWishlistInfo();

      info.activeItemsInfo.forEach((itemInfo, index) => {
        checkItemInfo(itemInfo, index, itemPrices[index]);
      });
    });

    it("should get exception if signature expired", async () => {
      const wishlistId = 123;
      const itemPrices = [wei(100, 6), wei(120, 6)];

      const sigDeadline = (await time.latest()) + 1000;

      const signature = await getCreateWishlistSignature(
        PROTOCOL_SIGNER,
        await FIRST.getAddress(),
        wishlistId,
        getItemPricesHash(itemPrices),
        sigDeadline,
        await wishlistFactory.getAddress(),
      );

      await time.setNextBlockTimestamp(sigDeadline + 10);

      await expect(
        wishlistFactory.connect(PROTOCOL_SIGNER).createWishlist(FIRST, wishlistId, sigDeadline, itemPrices, signature),
      ).to.be.revertedWithCustomError(wishlistFactory, "CreateWishlistSignatureExpired");
    });

    it("should get exception if wishlist with the same ID already exists", async () => {
      const wishlistId = 123;
      const itemPrices = [wei(100, 6), wei(120, 6)];

      const sigDeadline = (await time.latest()) + 1000;

      const signature = await getCreateWishlistSignature(
        PROTOCOL_SIGNER,
        await FIRST.getAddress(),
        wishlistId,
        getItemPricesHash(itemPrices),
        sigDeadline,
        await wishlistFactory.getAddress(),
      );

      await wishlistFactory
        .connect(PROTOCOL_SIGNER)
        .createWishlist(FIRST, wishlistId, sigDeadline, itemPrices, signature);

      await expect(
        wishlistFactory.connect(FIRST).createWishlist(FIRST, wishlistId, sigDeadline, itemPrices, signature),
      ).to.be.revertedWithCustomError(wishlistFactory, "WishlistAlreadyExists");
    });

    it("should get exception if pass invalid signature", async () => {
      const wishlistId = 123;
      const itemPrices = [wei(100, 6), wei(120, 6)];

      const sigDeadline = (await time.latest()) + 1000;

      const signature = await getCreateWishlistSignature(
        PROTOCOL_SIGNER,
        await FIRST.getAddress(),
        wishlistId,
        getItemPricesHash(itemPrices),
        sigDeadline,
        await wishlistFactory.getAddress(),
      );

      await expect(
        wishlistFactory.connect(FIRST).createWishlist(FIRST, 12, sigDeadline, itemPrices, signature),
      ).to.be.revertedWithCustomError(wishlistFactory, "InvalidCreateWishlistSignature");
    });
  });

  // describe("withdrawProtocolFee", () => {
  //   const feeAmount = wei(100, 6);

  //   beforeEach("setup", async () => {
  //     await usdcToken.transfer(wishlistFactory, feeAmount);
  //   });

  //   it("should correctly withdraw protocol fee from the contract", async () => {
  //     const tx = await wishlistFactory.withdrawProtocolFee(FEE_RECIPIENT);

  //     expect(tx).to.emit(wishlistFactory, "ProtocolFeeWithdrawn").withArgs(feeAmount, FEE_RECIPIENT);
  //     expect(await usdcToken.balanceOf(FEE_RECIPIENT)).to.be.eq(feeAmount);

  //     expect(await usdcToken.balanceOf(wishlistFactory)).to.be.eq(0);
  //   });

  //   it("should get exception if pass zero recipient address", async () => {
  //     await expect(wishlistFactory.withdrawProtocolFee(ethers.ZeroAddress))
  //       .to.revertedWithCustomError(wishlistFactory, "ZeroAddr")
  //       .withArgs("FeeRecipient");
  //   });

  //   it("should get exception if nothing to withdraw", async () => {
  //     await wishlistFactory.withdrawProtocolFee(FEE_RECIPIENT);

  //     await expect(wishlistFactory.withdrawProtocolFee(FEE_RECIPIENT))
  //       .to.revertedWithCustomError(wishlistFactory, "ZeroFeeToWithdraw")
  //       .withArgs();
  //   });

  //   it("should get exception if non-owner try to withdraw protocol fee", async () => {
  //     await expect(wishlistFactory.connect(FIRST).withdrawProtocolFee(FIRST)).to.revertedWith(
  //       "Ownable: caller is not the owner",
  //     );
  //   });
  // });

  describe("getWishlistAddresses", () => {
    it("should return correct wishlist addresses", async () => {
      const wishlistIds = [123, 124, 125];
      const signers = [OWNER, FIRST, SECOND];
      const txs = [];

      const sigDeadline = (await time.latest()) + 1000;

      for (let i = 0; i < wishlistIds.length; i++) {
        const signature = await getCreateWishlistSignature(
          PROTOCOL_SIGNER,
          await signers[i].getAddress(),
          wishlistIds[i],
          getItemPricesHash([]),
          sigDeadline,
          await wishlistFactory.getAddress(),
        );

        txs.push(
          await wishlistFactory
            .connect(signers[i])
            .createWishlist(signers[i], wishlistIds[i], sigDeadline, [], signature),
        );
      }

      expect(await wishlistFactory.getWishlistsTotalCount()).to.be.eq(wishlistIds.length);

      const wishlistAddresses = await wishlistFactory.getWishlistAddresses(0, 10);

      for (let i = 0; i < wishlistIds.length; i++) {
        await expect(txs[i])
          .to.emit(wishlistFactory, "NewWishlistCreated")
          .withArgs(await signers[i].getAddress(), wishlistIds[i], wishlistAddresses[i]);
      }
    });
  });
});
