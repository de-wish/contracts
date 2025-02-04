import { Deployer, Reporter } from "@solarity/hardhat-migrate";

import { getConfig } from "./config/config";

import {
  ERC20Mock__factory,
  ProtocolTreasury__factory,
  Wishlist__factory,
  WishlistFactory__factory,
  ERC1967Proxy__factory,
} from "@ethers-v6";

export = async (deployer: Deployer) => {
  const config = (await getConfig())!;

  let usdcTokenAddr: string = config.usdcTokenAddr;
  let signerAddr: string = await (await deployer.getSigner()).getAddress();

  if (!usdcTokenAddr) {
    const usdcToken = await deployer.deploy(ERC20Mock__factory, ["Test USDC", "USDC", 6]);

    usdcTokenAddr = await usdcToken.getAddress();
  }

  let contractsInfo: [string, string][] = [];

  const wishlistFactoryImpl = await deployer.deploy(WishlistFactory__factory);
  const wishlistImpl = await deployer.deploy(Wishlist__factory);

  const wishlistFactoryProxy = await deployer.deploy(
    ERC1967Proxy__factory,
    [await wishlistFactoryImpl.getAddress(), "0x"],
    { name: "WishlistFactoryProxy" },
  );

  const wishlistFactory = await deployer.deployed(WishlistFactory__factory, await wishlistFactoryProxy.getAddress());

  if (!config.protocolFeeSettings.protocolFeeRecipient) {
    const protocolTreasuryImpl = await deployer.deploy(ProtocolTreasury__factory);
    const protocolTreasuryProxy = await deployer.deploy(
      ERC1967Proxy__factory,
      [await protocolTreasuryImpl.getAddress(), "0x"],
      { name: "ProtocolTreasuryProxy" },
    );

    const protocolTreasury = await deployer.deployed(
      ProtocolTreasury__factory,
      await protocolTreasuryProxy.getAddress(),
    );

    await protocolTreasury.initialize();

    config.protocolFeeSettings.protocolFeeRecipient = await protocolTreasury.getAddress();

    contractsInfo.push(["ProtocolTreasury", await protocolTreasury.getAddress()]);
  }

  await wishlistFactory.initialize(
    await wishlistImpl.getAddress(),
    usdcTokenAddr,
    config.protocolSignerAddr ? config.protocolSignerAddr : signerAddr,
    config.protocolFeeSettings,
  );

  contractsInfo.push(
    ["USDC", usdcTokenAddr],
    ["WishlistFactory", await wishlistFactory.getAddress()],
    ["WishlistImpl", await wishlistImpl.getAddress()],
  );

  Reporter.reportContracts(...contractsInfo);
};
