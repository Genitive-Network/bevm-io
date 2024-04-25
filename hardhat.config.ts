import "dotenv/config";
import "hardhat-deploy";
import "@nomicfoundation/hardhat-toolbox";
import "@openzeppelin/hardhat-upgrades";

export default {
  namedAccounts: {
    deployer: 0, //the second account
  },
  networks: {
    testnet: {
      live: true,
      chainId: 97,
      url: process.env.RPC_TESTNET_URL,
      accounts: process.env.TESTNET_PRIVATE_KEY ? [`0x${process.env.TESTNET_PRIVATE_KEY}`] : [],
    },
    bevmtest: {
      live: true,
      chainId: 11503,
      url: process.env.BEVM_DEV_RPC_URL,
      accounts: process.env.BEVM_DEV_PRIVATE_KEY ? [`0x${process.env.BEVM_DEV_PRIVATE_KEY}`] : [],
    }
  },
  sourcify: {
    enabled: false,
    // Optional: specify a different Sourcify server
    apiUrl: "https://sourcify.dev/server",
    // Optional: specify a different Sourcify repository
    browserUrl: "https://repo.sourcify.dev",
  },
  solidity: {
    compilers: [
      {
        version: "0.8.20",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          }
        }
      },
      {
        version: "0.8.10",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          }
        }
      }
    ],
  },
  etherscan: {
    apiKey: {
      testnet: process.env.TEST_API_KEY,
      bevmtest: process.env.BEVM_API_KEY,
    },
    customChains: [
      {
        network: "testnet",
        chainId: 97, // Replace with the correct chainId for the "opbnb" network
        urls: {
          apiURL: "https://api-testnet.bscscan.com/api",
          browserURL: "https://testnet.bscscan.com/",
        },
      },
      {
        network: "bevmtest",
        chainId: 11503,
        urls: {
          apiURL: "https://api-testnet.polygonscan.com/api",
          browserURL: "https://scan-testnet.bevm.io"
        }
      },
    ],
  },
  paths: {
    sources: "./contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts"
  },
  mocha: {
    timeout: 40000
  }
};
