import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { BigNumberish, hexlify } from "ethers";
import { ethers } from "hardhat";

export function getItemPricesHash(itemPrices: BigNumberish[]) {
  return ethers.keccak256(ethers.AbiCoder.defaultAbiCoder().encode(["uint256[]"], [itemPrices]));
}

export async function getCreateWishlistSignature(
  sender: SignerWithAddress,
  wishlistOwner: string,
  wishlistId: BigNumberish,
  itemPricesHash: string,
  deadline: BigNumberish,
  myContractAddress: string,
  permitConfig?: { name?: string; chainId?: number; version?: string },
) {
  const [name, version, chainId] = await Promise.all([
    permitConfig?.name ?? "WishlistFactory",
    permitConfig?.version ?? "1",
    permitConfig?.chainId ?? (await sender.provider.getNetwork()).chainId,
  ]);

  const result = await sender.signTypedData(
    {
      name,
      version,
      chainId,
      verifyingContract: myContractAddress,
    },
    {
      CreateWishlist: [
        { name: "wishlistOwner", type: "address" },
        { name: "wishlistId", type: "uint256" },
        { name: "itemPricesHash", type: "bytes32" },
        { name: "deadline", type: "uint256" },
      ],
    },
    {
      wishlistOwner,
      wishlistId,
      itemPricesHash,
      deadline,
    },
  );

  return hexlify(result);
}
