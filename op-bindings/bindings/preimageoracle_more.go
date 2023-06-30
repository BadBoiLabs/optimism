// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PreimageOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageLengths\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1001,\"contract\":\"contracts/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageParts\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1002,\"contract\":\"contracts/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimagePartOk\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_mapping(t_uint256,t_bytes32)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bytes32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var PreimageOracleStorageLayout = new(solc.StorageLayout)

var PreimageOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106100725760003560e01c8063e159261111610050578063e15926111461011b578063fe4ac08e14610130578063fef2b4ed146101a557600080fd5b806361238bde146100775780638542cf50146100b5578063e03110e1146100f3575b600080fd5b6100a26100853660046103b6565b600160209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b6100e36100c33660046103b6565b600260209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016100ac565b6101066101013660046103b6565b6101c5565b604080519283526020830191909152016100ac565b61012e6101293660046103d8565b6102b6565b005b61012e61013e366004610454565b6000838152600260209081526040808320878452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558684528252808320968352958152858220939093559283529082905291902055565b6100a26101b3366004610486565b60006020819052908152604090205481565b6000828152600260209081526040808320848452909152812054819060ff1661024e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f707265696d616765206d75737420657869737400000000000000000000000000604482015260640160405180910390fd5b506000838152602081815260409091205461026a8160086104ce565b6102758560206104ce565b1061029357836102868260086104ce565b61029091906104e6565b91505b506000938452600160209081526040808620948652939052919092205492909150565b604435600080600883018611156102cc57600080fd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b600080604083850312156103c957600080fd5b50508035926020909101359150565b6000806000604084860312156103ed57600080fd5b83359250602084013567ffffffffffffffff8082111561040c57600080fd5b818601915086601f83011261042057600080fd5b81358181111561042f57600080fd5b87602082850101111561044157600080fd5b6020830194508093505050509250925092565b6000806000806080858703121561046a57600080fd5b5050823594602084013594506040840135936060013592509050565b60006020828403121561049857600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156104e1576104e161049f565b500190565b6000828210156104f8576104f861049f565b50039056fea164736f6c634300080f000a"

var PreimageOracleDeployedSourceMap = "57:2936:58:-:0;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;143:68;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;413:25:234;;;401:2;386:18;143:68:58;;;;;;;;217:66;;;;;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;614:14:234;;607:22;589:41;;577:2;562:18;217:66:58;449:187:234;290:454:58;;;;;;:::i;:::-;;:::i;:::-;;;;815:25:234;;;871:2;856:18;;849:34;;;;788:18;290:454:58;641:248:234;1537:1454:58;;;;;;:::i;:::-;;:::i;:::-;;1086:262;;;;;;:::i;:::-;1219:19;;;;:14;:19;;;;;;;;:31;;;;;;;;:38;;;;1253:4;1219:38;;;;;;1267:18;;;;;;;;:30;;;;;;;;;:37;;;;1314:20;;;;;;;;;;:27;1086:262;87:50;;;;;;:::i;:::-;;;;;;;;;;;;;;;290:454;388:11;439:19;;;:14;:19;;;;;;;;:27;;;;;;;;;388:11;;439:27;;431:59;;;;;;;2517:2:234;431:59:58;;;2499:21:234;2556:2;2536:18;;;2529:30;2595:21;2575:18;;;2568:49;2634:18;;431:59:58;;;;;;;;-1:-1:-1;521:14:58;538:20;;;509:2;538:20;;;;;;;;631:10;538:20;640:1;631:10;:::i;:::-;616:11;:6;625:2;616:11;:::i;:::-;:25;612:84;;679:6;666:10;:6;675:1;666:10;:::i;:::-;:19;;;;:::i;:::-;657:28;;612:84;-1:-1:-1;711:18:58;;;;:13;:18;;;;;;;;:26;;;;;;;;;;;;290:454;;-1:-1:-1;290:454:58:o;1537:1454::-;1831:4;1818:18;1636:12;;1958:1;1948:12;;1933:28;;1930:76;;;1990:1;1987;1980:12;1930:76;2249:3;2245:14;;;2149:4;2233:27;2280:11;2254:4;2399:15;2280:11;2381:40;2611:28;;;2615:11;2611:28;2605:35;2662:20;;;;2809:19;2802:27;2831:11;2799:44;2862:19;;;;2840:1;2862:19;;;;;;;;:31;;;;;;;;:38;;;;2896:4;2862:38;;;;;;2910:18;;;;;;;;:30;;;;;;;;;:37;;;;2957:20;;;;;;;;;;;:27;;;;-1:-1:-1;;;;1537:1454:58:o;14:248:234:-;82:6;90;143:2;131:9;122:7;118:23;114:32;111:52;;;159:1;156;149:12;111:52;-1:-1:-1;;182:23:234;;;252:2;237:18;;;224:32;;-1:-1:-1;14:248:234:o;894:659::-;973:6;981;989;1042:2;1030:9;1021:7;1017:23;1013:32;1010:52;;;1058:1;1055;1048:12;1010:52;1094:9;1081:23;1071:33;;1155:2;1144:9;1140:18;1127:32;1178:18;1219:2;1211:6;1208:14;1205:34;;;1235:1;1232;1225:12;1205:34;1273:6;1262:9;1258:22;1248:32;;1318:7;1311:4;1307:2;1303:13;1299:27;1289:55;;1340:1;1337;1330:12;1289:55;1380:2;1367:16;1406:2;1398:6;1395:14;1392:34;;;1422:1;1419;1412:12;1392:34;1467:7;1462:2;1453:6;1449:2;1445:15;1441:24;1438:37;1435:57;;;1488:1;1485;1478:12;1435:57;1519:2;1515;1511:11;1501:21;;1541:6;1531:16;;;;;894:659;;;;;:::o;1558:385::-;1644:6;1652;1660;1668;1721:3;1709:9;1700:7;1696:23;1692:33;1689:53;;;1738:1;1735;1728:12;1689:53;-1:-1:-1;;1761:23:234;;;1831:2;1816:18;;1803:32;;-1:-1:-1;1882:2:234;1867:18;;1854:32;;1933:2;1918:18;1905:32;;-1:-1:-1;1558:385:234;-1:-1:-1;1558:385:234:o;1948:180::-;2007:6;2060:2;2048:9;2039:7;2035:23;2031:32;2028:52;;;2076:1;2073;2066:12;2028:52;-1:-1:-1;2099:23:234;;1948:180;-1:-1:-1;1948:180:234:o;2663:184::-;2715:77;2712:1;2705:88;2812:4;2809:1;2802:15;2836:4;2833:1;2826:15;2852:128;2892:3;2923:1;2919:6;2916:1;2913:13;2910:39;;;2929:18;;:::i;:::-;-1:-1:-1;2965:9:234;;2852:128::o;2985:125::-;3025:4;3053:1;3050;3047:8;3044:34;;;3058:18;;:::i;:::-;-1:-1:-1;3095:9:234;;2985:125::o"

func init() {
	if err := json.Unmarshal([]byte(PreimageOracleStorageLayoutJSON), PreimageOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["PreimageOracle"] = PreimageOracleStorageLayout
	deployedBytecodes["PreimageOracle"] = PreimageOracleDeployedBin
}
