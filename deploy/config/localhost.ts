import { PRECISION, wei } from "@/scripts";

import { DeployConfig } from "./types";

export const deployConfig: DeployConfig = {
  usdcTokenAddr: "",
  protocolSignerAddr: "",
  protocolFeeSettings: {
    protocolFeePercentage: 2n * PRECISION,
    maxProtocolFeeAmount: wei(100, 6),
    protocolFeeRecipient: "",
  },
};
