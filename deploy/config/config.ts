import hre from "hardhat";

export async function getConfig() {
  if (hre.network.name == "localhost" || hre.network.name == "hardhat") {
    return (await import("./localhost")).deployConfig;
  }

  if (hre.network.name == "sepolia") {
    return (await import("./sepolia")).deployConfig;
  }

  throw new Error(`Config for network ${hre.network.name} is not specified`);
}
