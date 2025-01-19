import { ethers } from "hardhat";
import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { expect } from "chai";
import { Reverter } from "@test-helpers";
import { PRECISION, PERCENTAGE_100, wei } from "@scripts";
import { WishlistManager, IWishlistManager, ERC20Mock } from "@ethers-v6";
import { BigNumberish } from "ethers";

describe("WishlistManager", () => {
  const reverter = new Reverter();

  const initialTokensAmount = wei(10000);
  const defaultProtocolFee = 2n * PRECISION;
  const maxProtocolFeeAmount = wei(100, 6);

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let FEE_RECIPIENT: SignerWithAddress;

  let usdcToken: ERC20Mock;

  let wishlistManager: WishlistManager;

  function countFee(itemPrice: BigNumberish) {
    const protocolFee = (BigInt(itemPrice) * defaultProtocolFee) / PERCENTAGE_100;

    return protocolFee > maxProtocolFeeAmount ? maxProtocolFeeAmount : protocolFee;
  }

  function checkItemInfo(
    itemInfo: IWishlistManager.WishlistItemInfoStruct,
    expectedItemId: BigNumberish,
    expectedPrice: BigNumberish,
    expectedBuyerAddr: BigNumberish = ethers.ZeroAddress,
  ) {
    expect(itemInfo.itemId).to.be.eq(expectedItemId);
    expect(itemInfo.itemPrice).to.be.eq(expectedPrice);
    expect(itemInfo.itemBuyingFeeAmount).to.be.eq(countFee(expectedPrice));
    expect(itemInfo.buyerAddr).to.be.eq(expectedBuyerAddr);
    expect(itemInfo.isActive).to.be.eq(expectedBuyerAddr === ethers.ZeroAddress);
  }

  before(async () => {
    [OWNER, FIRST, SECOND, FEE_RECIPIENT] = await ethers.getSigners();

    const WishlistManagerFactory = await ethers.getContractFactory("WishlistManager");
    const ERC20MockFactory = await ethers.getContractFactory("ERC20Mock");

    usdcToken = await ERC20MockFactory.deploy("Mock USDC", "USDC", 6);

    await usdcToken.mint(OWNER, initialTokensAmount);
    await usdcToken.mint(FIRST, initialTokensAmount);

    wishlistManager = await WishlistManagerFactory.deploy(OWNER, usdcToken, defaultProtocolFee, maxProtocolFeeAmount);

    await usdcToken.approve(wishlistManager, initialTokensAmount);
    await usdcToken.connect(FIRST).approve(wishlistManager, initialTokensAmount);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("constructor", () => {
    it("should set parameters correctly", async () => {
      expect(await wishlistManager.owner()).to.eq(OWNER);
      expect(await wishlistManager.usdcToken()).to.eq(usdcToken);
      expect(await wishlistManager.protocolFeePercentage()).to.eq(defaultProtocolFee);
      expect(await wishlistManager.maxProtocolFeeAmount()).to.eq(maxProtocolFeeAmount);
    });
  });

  describe("setProtocolFeePercentage", () => {
    it("should correctly set new protocol fee value", async () => {
      const newProtocolFee = 3n * PRECISION;

      const tx = await wishlistManager.setProtocolFeePercentage(newProtocolFee);

      expect(tx).to.emit(wishlistManager, "ProtocolFeeUpdated").withArgs(newProtocolFee, defaultProtocolFee);

      expect(await wishlistManager.protocolFeePercentage()).to.be.eq(newProtocolFee);
    });

    it("should get exception if try to set protocol fee higher than the max value", async () => {
      const newProtocolFee = 110n * PRECISION;

      await expect(wishlistManager.setProtocolFeePercentage(newProtocolFee))
        .to.revertedWithCustomError(wishlistManager, "InvalidProtocolFee")
        .withArgs(newProtocolFee);
    });

    it("should get exception if non-owner try to set protocol fee", async () => {
      const newProtocolFee = 3n * PRECISION;

      await expect(wishlistManager.connect(FIRST).setProtocolFeePercentage(newProtocolFee))
        .to.revertedWithCustomError(wishlistManager, "OwnableUnauthorizedAccount")
        .withArgs(FIRST);
    });
  });

  describe("withdrawProtocolFee", () => {
    const wishlistId = 1;
    const initialItems = [wei(20, 6), wei(100, 6), wei(350, 6)];

    beforeEach("setup", async () => {
      await wishlistManager.connect(FIRST).createWishlist(initialItems);
    });

    it("should correctly withdraw protocol fee from the contract", async () => {
      await wishlistManager.buyWishlistItem(wishlistId, 0);
      await wishlistManager.buyWishlistItem(wishlistId, 1);

      const expectedFeeAmount = countFee(initialItems[0]) + countFee(initialItems[1]);

      expect(await wishlistManager.owedProtocolFee()).to.be.eq(expectedFeeAmount);

      const tx = await wishlistManager.withdrawProtocolFee(FEE_RECIPIENT);

      expect(tx).to.emit(wishlistManager, "ProtocolFeeWithdrawn").withArgs(expectedFeeAmount, FEE_RECIPIENT);
      expect(await usdcToken.balanceOf(FEE_RECIPIENT)).to.be.eq(expectedFeeAmount);

      expect(await wishlistManager.owedProtocolFee()).to.be.eq(0);
    });

    it("should get exception if pass zero recipient address", async () => {
      await wishlistManager.buyWishlistItem(wishlistId, 0);

      await expect(wishlistManager.withdrawProtocolFee(ethers.ZeroAddress))
        .to.revertedWithCustomError(wishlistManager, "ZeroFeeRecipient")
        .withArgs();
    });

    it("should get exception if nothing to withdraw", async () => {
      await expect(wishlistManager.withdrawProtocolFee(FEE_RECIPIENT))
        .to.revertedWithCustomError(wishlistManager, "ZeroFeeToWithdraw")
        .withArgs();
    });

    it("should get exception if non-owner try to withdraw protocol fee", async () => {
      await wishlistManager.buyWishlistItem(wishlistId, 0);

      await expect(wishlistManager.connect(FIRST).withdrawProtocolFee(FIRST))
        .to.revertedWithCustomError(wishlistManager, "OwnableUnauthorizedAccount")
        .withArgs(FIRST);
    });
  });

  describe("createWishlist", () => {
    it("should correctly create new wishlist without items", async () => {
      const expectedWishlistId = 1;

      const tx = await wishlistManager.connect(FIRST).createWishlist([]);

      expect(tx).to.emit(wishlistManager, "WishlistCreated").withArgs(FIRST, expectedWishlistId);

      expect(await wishlistManager.getUserWishlistIds(FIRST)).to.be.deep.eq([expectedWishlistId]);

      const info = await wishlistManager.getWishlistInfo(expectedWishlistId);

      expect(info.ownerAddr).to.be.eq(FIRST);
      expect(info.owedUsdcAmount).to.be.eq(0);
      expect(info.activeItemsInfo.length).to.be.eq(0);
    });

    it("should correctly create new wishlist with items", async () => {
      const expectedWishlistId = 1;
      const itemPrices = [wei(20, 6), wei(100, 6), wei(350, 6)];

      const tx = await wishlistManager.connect(FIRST).createWishlist(itemPrices);

      expect(tx).to.emit(wishlistManager, "WishlistCreated").withArgs(FIRST, expectedWishlistId);
      expect(tx).to.emit(wishlistManager, "WishlistItemsAdded").withArgs(expectedWishlistId, itemPrices.length);

      const info = await wishlistManager.getWishlistInfo(expectedWishlistId);

      expect(info.ownerAddr).to.be.eq(FIRST);
      expect(info.owedUsdcAmount).to.be.eq(0);
      expect(info.activeItemsInfo.length).to.be.eq(itemPrices.length);

      info.activeItemsInfo.forEach((itemInfo, index) => {
        checkItemInfo(itemInfo, index, itemPrices[index]);
      });
    });
  });

  describe("addWishlistItems", () => {
    const wishlistId = 1;
    const initialItems = [wei(100, 6)];

    beforeEach("setup", async () => {
      await wishlistManager.connect(FIRST).createWishlist(initialItems);
    });

    it("should correctly add new items", async () => {
      const itemPrices = [wei(20, 6), wei(100, 6), wei(350, 6)];

      const tx = await wishlistManager.connect(FIRST).addWishlistItems(wishlistId, itemPrices);

      expect(tx)
        .to.emit(wishlistManager, "WishlistItemsAdded")
        .withArgs(wishlistId, initialItems.length + itemPrices.length);

      const info = await wishlistManager.getWishlistInfo(wishlistId);

      const expectedItems = [...initialItems, ...itemPrices];

      info.activeItemsInfo.forEach((itemInfo, index) => {
        checkItemInfo(itemInfo, index, expectedItems[index]);
      });
    });

    it("should get exception if try to add empty items arr", async () => {
      await expect(wishlistManager.connect(FIRST).addWishlistItems(wishlistId, [])).to.revertedWithCustomError(
        wishlistManager,
        "ZeroItemsToAdd",
      );
    });

    it("should get exception if not a wishlist owner try to add new items", async () => {
      await expect(wishlistManager.connect(SECOND).addWishlistItems(wishlistId, [wei(100, 6)]))
        .to.revertedWithCustomError(wishlistManager, "NotAWishlistOwner")
        .withArgs(wishlistId, SECOND);
    });
  });

  describe("removeWishlistItems", () => {
    const wishlistId = 1;
    const initialItems = [wei(20, 6), wei(100, 6), wei(350, 6)];

    beforeEach("setup", async () => {
      await wishlistManager.connect(FIRST).createWishlist(initialItems);
    });

    it("should correctly remove items", async () => {
      const itemsToRemove = [1];

      const tx = await wishlistManager.connect(FIRST).removeWishlistItems(wishlistId, itemsToRemove);

      expect(tx).to.emit(wishlistManager, "WishlistItemsRemoved").withArgs(wishlistId, itemsToRemove);

      expect(await wishlistManager.isItemActive(wishlistId, itemsToRemove[0])).to.be.false;
    });

    it("should get exception if try to remove non-active item", async () => {
      const itemToRemove = 1;

      await wishlistManager.buyWishlistItem(wishlistId, itemToRemove);

      expect(await wishlistManager.isItemActive(wishlistId, itemToRemove)).to.be.false;

      await expect(wishlistManager.connect(FIRST).removeWishlistItems(wishlistId, [itemToRemove]))
        .to.revertedWithCustomError(wishlistManager, "NotAnActiveItem")
        .withArgs(wishlistId, itemToRemove);
    });

    it("should get exception if not a wishlist owner try to remove items", async () => {
      await expect(wishlistManager.connect(SECOND).removeWishlistItems(wishlistId, [0, 1]))
        .to.revertedWithCustomError(wishlistManager, "NotAWishlistOwner")
        .withArgs(wishlistId, SECOND);
    });
  });

  describe("changeWishlistItemsPrice", () => {
    const wishlistId = 1;
    const initialItems = [wei(20, 6), wei(100, 6), wei(350, 6)];

    beforeEach("setup", async () => {
      await wishlistManager.connect(FIRST).createWishlist(initialItems);
    });

    it("should correctly change item price", async () => {
      const changeItemPriceData: IWishlistManager.ChangedItemPriceDataStruct = {
        itemId: 1,
        newItemPrice: wei(120, 6),
      };

      await wishlistManager.connect(FIRST).changeWishlistItemsPrice(wishlistId, [changeItemPriceData]);

      expect(await wishlistManager.isItemActive(wishlistId, changeItemPriceData.itemId)).to.be.true;

      const itemInfo = await wishlistManager.getWishlistItemsInfo(wishlistId, [changeItemPriceData.itemId]);

      expect(itemInfo[0].itemPrice).to.be.eq(changeItemPriceData.newItemPrice);
    });

    it("should get exception if try to change price for non-active item", async () => {
      const changeItemPriceData: IWishlistManager.ChangedItemPriceDataStruct = {
        itemId: 1,
        newItemPrice: wei(120, 6),
      };

      await wishlistManager.buyWishlistItem(wishlistId, changeItemPriceData.itemId);

      expect(await wishlistManager.isItemActive(wishlistId, changeItemPriceData.itemId)).to.be.false;

      await expect(wishlistManager.connect(FIRST).changeWishlistItemsPrice(wishlistId, [changeItemPriceData]))
        .to.revertedWithCustomError(wishlistManager, "NotAnActiveItem")
        .withArgs(wishlistId, changeItemPriceData.itemId);
    });

    it("should get exception if not a wishlist owner try to change item price", async () => {
      const changeItemPriceData: IWishlistManager.ChangedItemPriceDataStruct = {
        itemId: 1,
        newItemPrice: wei(120, 6),
      };

      await expect(wishlistManager.connect(SECOND).changeWishlistItemsPrice(wishlistId, [changeItemPriceData]))
        .to.revertedWithCustomError(wishlistManager, "NotAWishlistOwner")
        .withArgs(wishlistId, SECOND);
    });
  });

  describe("withdrawFunds", () => {
    const wishlistId = 1;
    const initialItems = [wei(20, 6), wei(100, 6), wei(350, 6)];

    beforeEach("setup", async () => {
      await wishlistManager.connect(FIRST).createWishlist(initialItems);
    });

    it("should correctly withdraw wishlist tokens", async () => {
      await wishlistManager.buyWishlistItem(wishlistId, 0);
      await wishlistManager.buyWishlistItem(wishlistId, 1);

      const expectedTokensAmount = initialItems[0] + initialItems[1];

      let wishlistInfo = await wishlistManager.getWishlistInfo(wishlistId);

      expect(wishlistInfo.owedUsdcAmount).to.be.eq(expectedTokensAmount);

      const tx = await wishlistManager.connect(FIRST).withdrawFunds(wishlistId, FEE_RECIPIENT);

      expect(tx)
        .to.emit(wishlistManager, "WishlistFundsWithdrawn")
        .withArgs(wishlistId, expectedTokensAmount, FEE_RECIPIENT);
      expect(await usdcToken.balanceOf(FEE_RECIPIENT)).to.be.eq(expectedTokensAmount);

      wishlistInfo = await wishlistManager.getWishlistInfo(wishlistId);

      expect(wishlistInfo.owedUsdcAmount).to.be.eq(0);
    });

    it("should get exception if not a wishlist owner try to remove items", async () => {
      await expect(wishlistManager.connect(SECOND).withdrawFunds(wishlistId, SECOND))
        .to.revertedWithCustomError(wishlistManager, "NotAWishlistOwner")
        .withArgs(wishlistId, SECOND);
    });
  });

  describe("buyWishlistItem", () => {
    const wishlistId = 1;
    const initialItems = [wei(20, 6), wei(100, 6), wei(350, 6)];

    beforeEach("setup", async () => {
      await wishlistManager.connect(FIRST).createWishlist(initialItems);
    });

    it("should correctly buy item", async () => {
      const itemToBuy = 1;

      const expectedFeeAmount = countFee(initialItems[itemToBuy]);

      const tx = await wishlistManager.buyWishlistItem(wishlistId, itemToBuy);

      expect(tx).to.emit(wishlistManager, "ItemBought").withArgs(wishlistId, itemToBuy, initialItems[itemToBuy], OWNER);

      expect(await wishlistManager.isItemActive(wishlistId, itemToBuy)).to.be.false;

      const info = await wishlistManager.getWishlistInfo(wishlistId);

      expect(info.owedUsdcAmount).to.be.eq(initialItems[itemToBuy]);
      expect(await wishlistManager.owedProtocolFee()).to.be.eq(expectedFeeAmount);

      expect(await usdcToken.balanceOf(OWNER)).to.be.eq(
        initialTokensAmount - initialItems[itemToBuy] - expectedFeeAmount,
      );
    });
  });
});
