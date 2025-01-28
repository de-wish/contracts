import { BigNumberish } from "ethers";

export type DeployConfig = {
  usdcTokenAddr: string;
  protocolSignerAddr: string;
  protocolFeeSettings: ProtocolFeeSettings;
};

export type ProtocolFeeSettings = {
  protocolFeePercentage: BigNumberish;
  maxProtocolFeeAmount: BigNumberish;
  protocolFeeRecipient: string;
};
