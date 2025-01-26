import { ethers } from "hardhat";
import { BigNumberish } from "ethers";
import { expect } from "chai";

import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { time } from "@nomicfoundation/hardhat-network-helpers";

import { PERCENTAGE_100, PRECISION, wei } from "@scripts";
import { getCreateWishlistSignature, getItemPricesHash, Reverter } from "@test-helpers";

import { WishlistFactory, Wishlist, ERC20Mock, IWishlist } from "@ethers-v6";

describe("Wishlist", () => {
  const reverter = new Reverter();

  const wishlistId = 123;
  const initialTokensAmount = wei(10000);
  const defaultProtocolFeePercentage = 2n * PRECISION;
  const maxProtocolFeeAmount = wei(100, 6);

  let sigDeadline: number;

  let OWNER: SignerWithAddress;
  let FIRST: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let PROTOCOL_SIGNER: SignerWithAddress;
  let FEE_RECIPIENT: SignerWithAddress;

  let usdcToken: ERC20Mock;

  let wishlistImpl: Wishlist;
  let wishlistFactoryImpl: WishlistFactory;

  let wishlistFactory: WishlistFactory;
  let wishlist: Wishlist;

  function countFee(itemPrice: BigNumberish) {
    const protocolFee = (BigInt(itemPrice) * defaultProtocolFeePercentage) / PERCENTAGE_100;

    return protocolFee > maxProtocolFeeAmount ? maxProtocolFeeAmount : protocolFee;
  }

  function checkItemInfo(
    itemInfo: IWishlist.WishlistItemInfoStruct,
    expectedItemId: BigNumberish,
    expectedPrice: BigNumberish,
    expectedCollectedAmount: BigNumberish = 0,
    expectedBuyerAddresses: BigNumberish[] = [],
  ) {
    expect(itemInfo.itemId).to.be.eq(expectedItemId);
    expect(itemInfo.itemPrice).to.be.eq(expectedPrice);
    expect(itemInfo.collectedTokensAmount).to.be.eq(expectedCollectedAmount);
    expect(itemInfo.buyersAddresses).to.be.deep.eq(expectedBuyerAddresses);
    expect(itemInfo.isActive).to.be.eq(expectedCollectedAmount !== expectedPrice);
  }

  before(async () => {
    [OWNER, FIRST, SECOND, PROTOCOL_SIGNER, FEE_RECIPIENT] = await ethers.getSigners();

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

    await wishlistFactory.initialize(
      wishlistImpl,
      usdcToken,
      PROTOCOL_SIGNER,
      defaultProtocolFeePercentage,
      maxProtocolFeeAmount,
    );

    sigDeadline = (await time.latest()) + 1000;

    const signature = await getCreateWishlistSignature(
      PROTOCOL_SIGNER,
      await FIRST.getAddress(),
      wishlistId,
      getItemPricesHash([]),
      sigDeadline,
      await wishlistFactory.getAddress(),
    );

    await wishlistFactory.connect(FIRST).createWishlist(FIRST, wishlistId, sigDeadline, [], signature);

    wishlist = await ethers.getContractAt("Wishlist", await wishlistFactory.getWishlistAddress(wishlistId));

    await usdcToken.approve(wishlist, initialTokensAmount);
    await usdcToken.connect(FIRST).approve(wishlist, initialTokensAmount);

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

    it("should correctly init wishlist with initial items", async () => {
      const newWishlistId = 321;
      const itemPrices = [wei(100, 6), wei(50, 6)];

      const signature = await getCreateWishlistSignature(
        PROTOCOL_SIGNER,
        await FIRST.getAddress(),
        newWishlistId,
        getItemPricesHash(itemPrices),
        sigDeadline,
        await wishlistFactory.getAddress(),
      );

      await wishlistFactory.connect(FIRST).createWishlist(FIRST, newWishlistId, sigDeadline, itemPrices, signature);

      const newWishlist = await ethers.getContractAt("Wishlist", await wishlistFactory.getWishlistAddress(wishlistId));

      const info = await newWishlist.getWishlistInfo();

      info.activeItemsInfo.forEach((itemInfo, index) => {
        checkItemInfo(itemInfo, index, itemPrices[index]);
      });
    });

    it("should get exception if try to call init function twice", async () => {
      await expect(wishlist.initializeWishlist(0, FIRST, [])).to.be.revertedWith(
        "Initializable: contract is already initialized",
      );
    });
  });

  describe("addWishlistItems", () => {
    it("should correctly add new items", async () => {
      const itemPrices = [wei(20, 6), wei(100, 6), wei(350, 6)];

      const tx = await wishlist.connect(FIRST).addWishlistItems(itemPrices);

      expect(tx).to.emit(wishlist, "WishlistItemsAdded").withArgs(itemPrices.length);

      const info = await wishlist.getWishlistInfo();

      info.activeItemsInfo.forEach((itemInfo, index) => {
        checkItemInfo(itemInfo, index, itemPrices[index]);
      });
    });

    it("should get exception if try to add empty items arr", async () => {
      await expect(wishlist.connect(FIRST).addWishlistItems([])).to.revertedWithCustomError(wishlist, "ZeroItemsToAdd");
    });

    it("should get exception if try to add item with zero price", async () => {
      await expect(wishlist.connect(FIRST).addWishlistItems([0])).to.revertedWithCustomError(wishlist, "ZeroItemPrice");
    });

    it("should get exception if not a wishlist owner try to add new items", async () => {
      await expect(wishlist.connect(SECOND).addWishlistItems([wei(100, 6)]))
        .to.revertedWithCustomError(wishlist, "NotAWishlistOwner")
        .withArgs(SECOND);
    });
  });

  describe("removeWishlistItems", () => {
    const itemPrices = [wei(20, 6), wei(100, 6), wei(350, 6)];

    beforeEach("setup", async () => {
      await wishlist.connect(FIRST).addWishlistItems(itemPrices);
    });

    it("should correctly remove items", async () => {
      const itemsToRemove = [1];

      const tx = await wishlist.connect(FIRST).removeWishlistItems(itemsToRemove);

      expect(tx).to.emit(wishlist, "WishlistItemsRemoved").withArgs(itemsToRemove);

      expect(await wishlist.isItemActive(itemsToRemove[0])).to.be.false;
    });

    it("should get exception if try to remove non-active item", async () => {
      const itemToRemove = 1;

      await wishlist.buyWishlistItem(itemToRemove, itemPrices[itemToRemove]);

      expect(await wishlist.isItemActive(itemToRemove)).to.be.false;

      await expect(wishlist.connect(FIRST).removeWishlistItems([itemToRemove]))
        .to.revertedWithCustomError(wishlist, "NotAnActiveItem")
        .withArgs(itemToRemove);
    });

    it("should get exception if not a wishlist owner try to remove items", async () => {
      await expect(wishlist.connect(SECOND).removeWishlistItems([0, 1]))
        .to.revertedWithCustomError(wishlist, "NotAWishlistOwner")
        .withArgs(SECOND);
    });
  });

  describe("changeWishlistItemsPrice", () => {
    const itemPrices = [wei(20, 6), wei(100, 6), wei(350, 6)];

    beforeEach("setup", async () => {
      await wishlist.connect(FIRST).addWishlistItems(itemPrices);
    });

    it("should correctly change item price", async () => {
      const changeItemPriceData: IWishlist.ChangedItemPriceDataStruct = {
        itemId: 1,
        newItemPrice: wei(120, 6),
      };

      await wishlist.connect(FIRST).changeWishlistItemsPrice([changeItemPriceData]);

      expect(await wishlist.isItemActive(changeItemPriceData.itemId)).to.be.true;

      const itemInfo = await wishlist.getWishlistItemsInfo([changeItemPriceData.itemId]);

      expect(itemInfo[0].itemPrice).to.be.eq(changeItemPriceData.newItemPrice);
    });

    it("should get exception if try to change price for non-active item", async () => {
      const itemId = 1;
      const changeItemPriceData: IWishlist.ChangedItemPriceDataStruct = {
        itemId,
        newItemPrice: wei(120, 6),
      };

      await wishlist.buyWishlistItem(changeItemPriceData.itemId, itemPrices[itemId]);

      expect(await wishlist.isItemActive(changeItemPriceData.itemId)).to.be.false;

      await expect(wishlist.connect(FIRST).changeWishlistItemsPrice([changeItemPriceData]))
        .to.revertedWithCustomError(wishlist, "NotAnActiveItem")
        .withArgs(changeItemPriceData.itemId);
    });

    it("should get exception if not a wishlist owner try to change item price", async () => {
      const changeItemPriceData: IWishlist.ChangedItemPriceDataStruct = {
        itemId: 1,
        newItemPrice: wei(120, 6),
      };

      await expect(wishlist.connect(SECOND).changeWishlistItemsPrice([changeItemPriceData]))
        .to.revertedWithCustomError(wishlist, "NotAWishlistOwner")
        .withArgs(SECOND);
    });
  });

  describe("withdrawFunds", () => {
    const itemPrices = [wei(20, 6), wei(100, 6), wei(350, 6)];

    beforeEach("setup", async () => {
      await wishlist.connect(FIRST).addWishlistItems(itemPrices);
    });

    it("should correctly withdraw wishlist tokens", async () => {
      await wishlist.buyWishlistItem(0, itemPrices[0]);
      await wishlist.buyWishlistItem(1, itemPrices[1]);

      const expectedTokensAmount = itemPrices[0] + itemPrices[1];

      let wishlistInfo = await wishlist.getWishlistInfo();

      expect(wishlistInfo.owedUsdcAmount).to.be.eq(expectedTokensAmount);

      const tx = await wishlist.connect(FIRST).withdrawFunds(FEE_RECIPIENT);

      expect(tx).to.emit(wishlist, "WishlistFundsWithdrawn").withArgs(expectedTokensAmount, FEE_RECIPIENT);
      expect(await usdcToken.balanceOf(FEE_RECIPIENT)).to.be.eq(expectedTokensAmount);

      wishlistInfo = await wishlist.getWishlistInfo();

      expect(wishlistInfo.owedUsdcAmount).to.be.eq(0);
    });

    it("should get exception if nothing to withdraw", async () => {
      await expect(wishlist.connect(FIRST).withdrawFunds(FIRST)).to.revertedWithCustomError(
        wishlist,
        "ZeroAmountToWithdraw",
      );
    });

    it("should get exception if not a wishlist owner try to remove items", async () => {
      await expect(wishlist.connect(SECOND).withdrawFunds(SECOND))
        .to.revertedWithCustomError(wishlist, "NotAWishlistOwner")
        .withArgs(SECOND);
    });
  });

  describe("buyWishlistItem", () => {
    const itemPrices = [wei(20, 6), wei(100, 6), wei(350, 6)];

    beforeEach("setup", async () => {
      await wishlist.connect(FIRST).addWishlistItems(itemPrices);
    });

    it("should correctly fully buy item", async () => {
      const itemToBuy = 1;

      const expectedFeeAmount = countFee(itemPrices[itemToBuy]);

      const tx = await wishlist.buyWishlistItem(itemToBuy, itemPrices[itemToBuy]);

      expect(tx)
        .to.emit(wishlist, "ItemBought")
        .withArgs(itemToBuy, itemPrices[itemToBuy], expectedFeeAmount, OWNER, true);

      expect(await wishlist.isItemActive(itemToBuy)).to.be.false;

      const info = await wishlist.getWishlistInfo();

      expect(info.owedUsdcAmount).to.be.eq(itemPrices[itemToBuy]);
      expect(await usdcToken.balanceOf(wishlistFactory)).to.be.eq(expectedFeeAmount);

      expect(await usdcToken.balanceOf(OWNER)).to.be.eq(
        initialTokensAmount - itemPrices[itemToBuy] - expectedFeeAmount,
      );
    });
  });
});
