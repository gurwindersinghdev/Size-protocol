// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./accounts/Receiver.sol";
import "./accounts/ERC1271.sol";
import "./accounts/ERC4337.sol";
import "./accounts/ERC4337Factory.sol";
import "./accounts/ERC6551.sol";
import "./accounts/ERC6551Proxy.sol";
import "./auth/Ownable.sol";
import "./auth/OwnableRoles.sol";
import "./tokens/WETH.sol";
import "./tokens/ERC20.sol";
import "./tokens/ERC4626.sol";
import "./tokens/ERC721.sol";
import "./tokens/ERC2981.sol";
import "./tokens/ERC1155.sol";
import "./tokens/ERC6909.sol";
import "./utils/Base64.sol";
import "./utils/CREATE3.sol";
import "./utils/Clone.sol";
import "./utils/DateTimeLib.sol";
import "./utils/DynamicBufferLib.sol";
import "./utils/ECDSA.sol";
import "./utils/EIP712.sol";
import "./utils/ERC1967Factory.sol";
import "./utils/ERC1967FactoryConstants.sol";
import "./utils/FixedPointMathLib.sol";
import "./utils/GasBurnerLib.sol";
import "./utils/Initializable.sol";
import "./utils/JSONParserLib.sol";
import "./utils/LibBit.sol";
import "./utils/LibBitmap.sol";
import "./utils/LibMap.sol";
import "./utils/LibClone.sol";
import "./utils/LibPRNG.sol";
import "./utils/LibRLP.sol";
import "./utils/LibSort.sol";
import "./utils/LibString.sol";
import "./utils/LibZip.sol";
import "./utils/MerkleProofLib.sol";
import "./utils/MetadataReaderLib.sol";
import "./utils/MinHeapLib.sol";
import "./utils/Multicallable.sol";
import "./utils/RedBlackTreeLib.sol";
import "./utils/SSTORE2.sol";
import "./utils/SafeCastLib.sol";
import "./utils/SafeTransferLib.sol";
import "./utils/SignatureCheckerLib.sol";
import "./utils/UUPSUpgradeable.sol";

library Milady {
    string internal constant WEBSITE = "https://miladymaker.net";

    address internal constant CONTRACT = 0x5Af0D9827E0c53E4799BB226655A1de152A425a5;
}
