# De-Wish Contracts

## Overview

This repository contains smart contracts for managing user wishlists, handling contributions, and managing protocol fees. The key contracts include `Wishlist` and `WishlistFactory`.

### Compilation

To compile the contracts, use the next script:

```bash
npm run compile
```

### Test

To run the tests, execute the following command:

```bash
npm run test
```

Or to see the coverage, run:

```bash
npm run coverage
```

### Deployment

To deploy smart contracts you need to do the following steps:
1. Create and fill in the `.env` file according to the template from the `.env.example` file
2. Fill in the required config file in the `deploy/config` directory according to the example below

```typescript
deployConfig = {
  usdcTokenAddr: "",
  protocolSignerAddr: "",
  protocolFeeSettings: {
    protocolFeePercentage: 2n * PRECISION,
    maxProtocolFeeAmount: wei(100, 6),
    protocolFeeRecipient: "",
  },
};
```

3. Run deploy command from the `package.json` file

```bash
npm run deploy-<network>
```

### Bindings

The command to generate the bindings is as follows:

```bash
npm run generate-types
```

> See the full list of available commands in the `package.json` file.