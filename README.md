# Capital47
Cross-Border Payments with Smart Contracts and CaptialOne APIs

## Background

Cross border payments are a pain. They are slow, expensive, and opaque. We want to make them fast, cheap, and transparent. Our solution enables to rethink the way we send money across borders.

### Streamlined Cross-Border Transactions with Blockchain and Cryptocurrencies

### 1. Traditional Cross-Border Transactions:
   - In traditional financial systems, cross-border transactions involve multiple intermediary banks, currency conversions, and can take several days to complete.
   - High fees are often associated with these transactions due to the involvement of various financial institutions and currency exchange processes.
   - Exchange rates may fluctuate, leading to uncertainty about the final amount received by the recipient.

### 2. Blockchain and Cryptocurrencies:
   - Blockchain is a decentralized and distributed ledger technology that records transactions across a network of computers in a secure and transparent manner.
   - Cryptocurrencies, such as Bitcoin or Ethereum, operate on blockchain technology. They are digital or virtual currencies that use cryptography for security and operate independently of a central authority.

### 3. Streamlining Cross-Border Transactions:
   - By leveraging blockchain and cryptocurrencies, cross-border transactions can be streamlined and made more efficient.
   - Cryptocurrencies can be transferred directly between users on a blockchain without the need for multiple intermediaries.
   - The decentralized nature of blockchain ensures transparency, reduces the risk of fraud, and provides a tamper-resistant record of transactions.

### 4. Benefits of Using Cryptocurrencies for Cross-Border Transactions:
   - **Speed:** Cryptocurrency transactions can be processed much faster compared to traditional bank transfers, which may take several days. Transactions on a blockchain can occur in near real-time.
   - **Cost-Effectiveness:** Cryptocurrency transactions often have lower fees compared to traditional cross-border transfers. This is because there are fewer intermediaries and less infrastructure involved.
   - **Reduced Currency Exchange Costs:** Cryptocurrencies are not tied to specific countries, reducing the need for currency conversions and associated fees.
   - **24/7 Accessibility:** Cryptocurrency transactions are not bound by banking hours or holidays, allowing users to send or receive funds at any time.

We are using `CapitalOne`'s `Nessie RESTful APIs` to integrate with the ACH network and EVM based `stablecoin smart contracts` to provide a transparent and secure way to send money across borders. The smart contracts are deployed on the `Hedera` testnet and the front end is deployed on Netlify.

## Why Hedera?

Hedera enables anyone to easily develop secure, fair, blazing-fast decentralized applications. It can be used to build stablecoins smart contracts, which are used in our project. Hedera smart contracts gas fees are very low, and it is independent of the transaction amount unlike traditional ACH network transactions.

Most importantly, Hedera is carbon-neutral and energy efficient, which is a huge plus for us.

## Demo

### Hedera Stablecoin Demo:

[Hedera Stablecoin Smart Contracts](Hedera.mov)


```js
{
  hederaTokenManager: '0.0.636684',
  name: 'C47Coin',
  symbol: 'C47',
  decimals: 6,
  initialSupply: undefined,
  supplyType: 'INFINITE',
  maxSupply: undefined,
  freezeKey: 'The Smart Contract',
  KYCKey: PublicKey {
    key: '08dda86fcdfd86ce467c517ed25175544b9e6b58bca319395fdb4f9cd0dce1b9',
    type: 'ED25519'
  },
  wipeKey: 'The Smart Contract',
  adminKey: 'The Smart Contract',
  supplyKey: 'The Smart Contract',
  pauseKey: 'The Smart Contract',
  feeScheduleKey: 'None',
  treasury: 'The Smart Contract',
  reserve: 'Proof of Reserve Feed initial amount : 1000',
  burnRole: '0.0.5904865',
  wipeRole: '0.0.5904865',
  rescueRole: '0.0.5904865',
  pauseRole: '0.0.5904865',
  freezeRole: '0.0.5904865',
  deleteRole: '0.0.5904865',
  kycRole: undefined,
  cashinRole: '0.0.5904865',
  cashinAllowance: '0',
  metadata: undefined,
  proxyAdminOwnerAccount: '0.0.5904865'
}
```

### Application Demo:

[Application Demo](application.mov)


## Technologies Used

- TypeScript
- Hedera Stablecoin Smart Contracts
- CapitalOne Nessie API
- React (for frontend)
- Go (for backend)

## Installation

## Smart Contracts

Smart contracts are deployed on the Hedera testnet.

### Backend

```bash
cd backend
go run main.go
```

### Frontend

```bash
cd web
npm install
npm start
```

## Acknowledgements

- [Hedera](https://hedera.com/)
- [CapitalOne Restful APIs](http://api.nessieisreal.com/)

## Contact

- [Rohith Reddy Kota](https://www.linkedin.com/in/rohithreddykota/)
- [Anusha Devi](https://www.linkedin.com/in/anushareddykota/)
- [Pramod Kumar](https://www.linkedin.com/in/pramod-kumar-undrakonda/)
- [Sravanthi Nittala](https://www.linkedin.com/in/sravanthi-nittala/)
