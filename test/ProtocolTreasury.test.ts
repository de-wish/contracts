import { ethers } from "hardhat";
import { expect } from "chai";

import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";

import { wei } from "@scripts";
import { Reverter } from "@test-helpers";

import { ERC20Mock, ProtocolTreasury } from "@ethers-v6";

describe("ProtocolTreasury", () => {
  const reverter = new Reverter();

  const initialTokensAmount = wei(10000);

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;

  let testToken: ERC20Mock;

  let protocolTreasury: ProtocolTreasury;

  before(async () => {
    [OWNER, FIRST] = await ethers.getSigners();

    const ERC1967ProxyFactory = await ethers.getContractFactory("ERC1967Proxy");
    const ProtocolTreasuryFactory = await ethers.getContractFactory("ProtocolTreasury");
    const ERC20MockFactory = await ethers.getContractFactory("ERC20Mock");

    testToken = await ERC20MockFactory.deploy("Mock USDC", "USDC", 6);

    await testToken.mint(OWNER, initialTokensAmount);
    await testToken.mint(FIRST, initialTokensAmount);

    const protocolTreasuryImpl = await ProtocolTreasuryFactory.deploy();

    const protocolTreasuryProxy = await ERC1967ProxyFactory.deploy(protocolTreasuryImpl, "0x");

    protocolTreasury = await ethers.getContractAt("ProtocolTreasury", protocolTreasuryProxy);

    await protocolTreasury.initialize();

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("initialize", () => {
    it("should set parameters correctly", async () => {
      expect(await protocolTreasury.owner()).to.eq(OWNER);
    });

    it("should get exception if try to call init function twice", async () => {
      await expect(protocolTreasury.initialize()).to.be.revertedWith("Initializable: contract is already initialized");
    });
  });

  describe("upgradability", () => {
    it("should correctly upgrade ProtocolTreasury contract", async () => {
      const ProtocolTreasuryMockFactory = await ethers.getContractFactory("ProtocolTreasuryMock");

      const wishlistFactoryMockImpl = await ProtocolTreasuryMockFactory.deploy();

      const newProtocolTreasury = await ethers.getContractAt("ProtocolTreasuryMock", protocolTreasury);

      await expect(newProtocolTreasury.version()).to.be.revertedWithoutReason();

      await protocolTreasury.upgradeTo(wishlistFactoryMockImpl);

      expect(await newProtocolTreasury.version()).to.be.eq("3.3.3");
    });

    it("should get exception if non-owner try to upgrade WishlistFactory contract", async () => {
      await expect(protocolTreasury.connect(FIRST).upgradeTo(ethers.ZeroAddress)).to.revertedWith(
        "Ownable: caller is not the owner",
      );
    });
  });

  describe("receive", () => {
    it("should correctly receive native currency", async () => {
      const nativeAmountToSend = wei(1);

      expect(await protocolTreasury.getSelfNativeBalance()).to.be.eq(0);

      await OWNER.sendTransaction({
        to: protocolTreasury,
        value: nativeAmountToSend,
      });

      expect(await protocolTreasury.getSelfNativeBalance()).to.be.eq(nativeAmountToSend);
    });
  });

  describe("withdrawNative", () => {
    const nativeAmountToSend = wei(1);
    const withdrawAmount = wei(0.5);

    beforeEach("setup", async () => {
      await OWNER.sendTransaction({
        to: protocolTreasury,
        value: nativeAmountToSend,
      });
    });

    it("should correctly withdraw native currency from the protocol treasury", async () => {
      const firstBalanceBefore = await ethers.provider.getBalance(FIRST);

      const tx = await protocolTreasury.withdrawNative(FIRST, withdrawAmount, false);

      await expect(tx).to.emit(protocolTreasury, "NativeCurrencyWithdrawn").withArgs(FIRST, withdrawAmount);

      expect(await protocolTreasury.getSelfNativeBalance()).to.be.eq(nativeAmountToSend - withdrawAmount);
      expect(await ethers.provider.getBalance(FIRST)).to.be.eq(firstBalanceBefore + withdrawAmount);
    });

    it("should correctly withdraw all native currency from the protocol treasury", async () => {
      const firstBalanceBefore = await ethers.provider.getBalance(FIRST);

      const tx = await protocolTreasury.withdrawNative(FIRST, 0, true);

      await expect(tx).to.emit(protocolTreasury, "NativeCurrencyWithdrawn").withArgs(FIRST, nativeAmountToSend);

      expect(await protocolTreasury.getSelfNativeBalance()).to.be.eq(0);
      expect(await ethers.provider.getBalance(FIRST)).to.be.eq(firstBalanceBefore + nativeAmountToSend);
    });

    it("should get exception if non-owner try to call this function", async () => {
      await expect(protocolTreasury.connect(FIRST).withdrawNative(FIRST, 0, true)).to.revertedWith(
        "Ownable: caller is not the owner",
      );
    });
  });

  describe("withdrawERC20Tokens", () => {
    const amountToSend = wei(1000);
    const withdrawAmount = wei(200);

    beforeEach("setup", async () => {
      await testToken.transfer(protocolTreasury, amountToSend);
    });

    it("should correctly withdraw native currency from the protocol treasury", async () => {
      const tx = await protocolTreasury.withdrawERC20Tokens(testToken, FIRST, withdrawAmount, false);

      await expect(tx).to.emit(protocolTreasury, "ERC20TokensWithdrawn").withArgs(testToken, FIRST, withdrawAmount);

      expect(await protocolTreasury.getSelfERC20TokenBalance(testToken)).to.be.eq(amountToSend - withdrawAmount);
      expect(await testToken.balanceOf(FIRST)).to.be.eq(initialTokensAmount + withdrawAmount);
    });

    it("should correctly withdraw all native currency from the protocol treasury", async () => {
      const tx = await protocolTreasury.withdrawERC20Tokens(testToken, FIRST, 0, true);

      await expect(tx).to.emit(protocolTreasury, "ERC20TokensWithdrawn").withArgs(testToken, FIRST, amountToSend);

      expect(await protocolTreasury.getSelfNativeBalance()).to.be.eq(0);
      expect(await testToken.balanceOf(FIRST)).to.be.eq(initialTokensAmount + amountToSend);
    });

    it("should get exception if non-owner try to call this function", async () => {
      await expect(protocolTreasury.connect(FIRST).withdrawERC20Tokens(testToken, FIRST, 0, true)).to.revertedWith(
        "Ownable: caller is not the owner",
      );
    });
  });
});
