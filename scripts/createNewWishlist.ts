import { Reporter } from "@solarity/hardhat-migrate";
import { ethers } from "hardhat";

import { wei } from "@scripts";

import { getCreateWishlistSignature, getItemPricesHash } from "@test-helpers";

async function main() {
  const wishlistFactoryAddr = "";

  const wishlistId = 123;
  const wishlistOwner = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266";
  const initialItems = [wei(100, 6), wei(50, 6)];
  const sigDeadline = Date.now() + 600;

  const signer = (await ethers.getSigners())[0];

  const signature = await getCreateWishlistSignature(
    signer,
    wishlistOwner,
    wishlistId,
    getItemPricesHash(initialItems),
    sigDeadline,
    wishlistFactoryAddr,
  );

  const wishlistFactory = await ethers.getContractAt("WishlistFactory", wishlistFactoryAddr);

  await wishlistFactory.createWishlist(wishlistOwner, wishlistId, sigDeadline, initialItems, signature);

  const wishlistAddr = await wishlistFactory.getWishlistAddress(wishlistId);

  Reporter.reportContracts(["DeployedWishlistAddr", wishlistAddr]);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});

// npx hardhat --network localhost run scripts/createNewWishlist.ts
