# Changelog

## [v0.26.0](https://github.com/mesg-foundation/engine/releases/tag/v0.26.0)

#### Breaking Changes

- ([#1842](https://github.com/mesg-foundation/engine/pull/1842)) Execution price based on service task price and new credit module.
- ([#1853](https://github.com/mesg-foundation/engine/pull/1853)) Fix 3 issues with instance.
- ([#1855](https://github.com/mesg-foundation/engine/pull/1855)) Add modules gov, upgrade, evidence and crisis.

#### Added

- ([#1860](https://github.com/mesg-foundation/engine/pull/1860)) Add export and import genesis functions to module credit .
- ([#1861](https://github.com/mesg-foundation/engine/pull/1861)) Add msg to mint credits module credit.

#### Changed

- ([#1849](https://github.com/mesg-foundation/engine/pull/1849)) Add a mechanism to reuse previous build image to speed up build time.
- ([#1850](https://github.com/mesg-foundation/engine/pull/1850)) Move CLI build and publish script to makefile to avoid duplication of code.

#### Fixed

- ([#1862](https://github.com/mesg-foundation/engine/pull/1862)) Improve checks of commands' parameters.

#### Removed

- ([#1863](https://github.com/mesg-foundation/engine/pull/1863)) Remove not used expected keeper.

## [v0.25.1](https://github.com/mesg-foundation/engine/releases/tag/v0.25.1)

#### Changed

- ([#1845](https://github.com/mesg-foundation/engine/pull/1845)) Set default flag gasPrice of the orchestrator to null.
- ([#1851](https://github.com/mesg-foundation/engine/pull/1851)) Remove swarm init from ci.

#### Fixed

- ([#1838](https://github.com/mesg-foundation/engine/pull/1838)) Remove min gas price from dev chain and e2e.
- ([#1841](https://github.com/mesg-foundation/engine/pull/1841)) Fix dev-starter when user stop script during the first sleep of 5 sec.
- ([#1843](https://github.com/mesg-foundation/engine/pull/1843)) Display body in lcd client on error.
- ([#1844](https://github.com/mesg-foundation/engine/pull/1844)) Add a check in LCD client if the account exists or not.
- ([#1846](https://github.com/mesg-foundation/engine/pull/1846)) Fix command execution list because of not set filter.
- ([#1848](https://github.com/mesg-foundation/engine/pull/1848)) Implement import and export functions in modules.

#### Dependencies

- ([#1835](https://github.com/mesg-foundation/engine/pull/1835)) Bump github.com/cosmos/cosmos-sdk from 0.38.3 to 0.38.4.

## [v0.25.0](https://github.com/mesg-foundation/engine/releases/tag/v0.25.0)

#### Breaking Changes

- ([#1830](https://github.com/mesg-foundation/engine/pull/1830)) Revert "Fix 2 issues with instance".

#### Added

- ([#1825](https://github.com/mesg-foundation/engine/pull/1825)) Add filter on execution list.

#### Changed

- ([#1827](https://github.com/mesg-foundation/engine/pull/1827)) Update dev script to use docker container instead of docker service.
- ([#1828](https://github.com/mesg-foundation/engine/pull/1828)) Improve CLI commands.
- ([#1831](https://github.com/mesg-foundation/engine/pull/1831)) Switch to json logger in orchestrator and daemon.

## [v0.24.0](https://github.com/mesg-foundation/engine/releases/tag/v0.24.0)

#### Breaking Changes

- ([#1810](https://github.com/mesg-foundation/engine/pull/1810)) Switch to CLI based Engine.

#### Added

- ([#1798](https://github.com/mesg-foundation/engine/pull/1798)) Add tx commands for service, process, execution and runner modules.
- ([#1801](https://github.com/mesg-foundation/engine/pull/1801)) Add orchestrator command to CLI.
- ([#1806](https://github.com/mesg-foundation/engine/pull/1806)) Create pre-configured dev docker image.
- ([#1808](https://github.com/mesg-foundation/engine/pull/1808)) Add cli dockerfile.
- ([#1812](https://github.com/mesg-foundation/engine/pull/1812)) Add logs to orchestrator.

#### Changed

- ([#1802](https://github.com/mesg-foundation/engine/pull/1802)) Rename in cosmos lcd and rpc clients minGasPrices by simply gasPrices.
- ([#1803](https://github.com/mesg-foundation/engine/pull/1803)) Update readme.
- ([#1817](https://github.com/mesg-foundation/engine/pull/1817)) Use docker container in e2e tests.

#### Fixed

- ([#1760](https://github.com/mesg-foundation/engine/pull/1760)) Set grace period to 1m in dev script.
- ([#1809](https://github.com/mesg-foundation/engine/pull/1809)) Fix runner e2e that block the tests on error.

#### Dependencies

- ([#1745](https://github.com/mesg-foundation/engine/pull/1745)) Bump github.com/sirupsen/logrus from 1.4.2 to 1.5.0.
- ([#1777](https://github.com/mesg-foundation/engine/pull/1777)) Bump github.com/spf13/viper from 1.6.2 to 1.6.3.
- ([#1782](https://github.com/mesg-foundation/engine/pull/1782)) Bump github.com/golang/protobuf from 1.3.4 to 1.4.0.
- ([#1794](https://github.com/mesg-foundation/engine/pull/1794)) Bump google.golang.org/grpc from 1.28.0 to 1.29.1.
- ([#1800](https://github.com/mesg-foundation/engine/pull/1800)) Bump github.com/prometheus/client_golang from 1.5.0 to 1.6.0.
- ([#1811](https://github.com/mesg-foundation/engine/pull/1811)) Revert golang/protobuf to v1.3.5 and prometheus/client_golang to v1.5.1.

## [v0.23.0](https://github.com/mesg-foundation/engine/releases/tag/v0.23.0)

#### Added

- ([#1795](https://github.com/mesg-foundation/engine/pull/1795)) Add lcd endpoint to estimate the gas of a transaction.

#### Fixed

- ([#1797](https://github.com/mesg-foundation/engine/pull/1797)) Fix 2 issues with instance.

#### Removed

- ([#1796](https://github.com/mesg-foundation/engine/pull/1796)) Remove ipfs endpoint from config.

## [v0.22.0](https://github.com/mesg-foundation/engine/releases/tag/v0.22.0)

#### Breaking Changes

- ([#1785](https://github.com/mesg-foundation/engine/pull/1785)) Replace bech32 prefix from "mesgtest" to "mesg".

## [v0.21.0](https://github.com/mesg-foundation/engine/releases/tag/v0.21.0)

#### Breaking Changes

- ([#1737](https://github.com/mesg-foundation/engine/pull/1737)) Add support for any type of data to filter.
- ([#1739](https://github.com/mesg-foundation/engine/pull/1739)) Add reference system to filter.
- ([#1774](https://github.com/mesg-foundation/engine/pull/1774)) Remove deprecated gRPC API.
- ([#1783](https://github.com/mesg-foundation/engine/pull/1783)) Calculate the instance env hash with only the customized env.

#### Added

- ([#1749](https://github.com/mesg-foundation/engine/pull/1749)) Add more validation on resources.
- ([#1758](https://github.com/mesg-foundation/engine/pull/1758)) Introduce a simple LCD client.
- ([#1759](https://github.com/mesg-foundation/engine/pull/1759)) add 2 more linters and remove useless exclusion.
- ([#1767](https://github.com/mesg-foundation/engine/pull/1767)) New Runner gRPC API.
- ([#1768](https://github.com/mesg-foundation/engine/pull/1768)) Add lcd endpoints and commands exists to runner and process.
- ([#1771](https://github.com/mesg-foundation/engine/pull/1771)) Add Orchestrator gRPC API.
- ([#1780](https://github.com/mesg-foundation/engine/pull/1780)) Add info to cosmos version on build.

#### Changed

- ([#1751](https://github.com/mesg-foundation/engine/pull/1751)) Replace gRPC by LCD as much as possible from e2e tests.
- ([#1754](https://github.com/mesg-foundation/engine/pull/1754)) Update lcd service hash endpoint to use a dedicated request structure.
- ([#1755](https://github.com/mesg-foundation/engine/pull/1755)) Update lcd process hash endpoint to use a dedicated request structure.
- ([#1756](https://github.com/mesg-foundation/engine/pull/1756)) Change gogoproto's customtype to casttype.
- ([#1757](https://github.com/mesg-foundation/engine/pull/1757)) Replace runner gRPC by LCD in e2e tests.
- ([#1761](https://github.com/mesg-foundation/engine/pull/1761)) Remove module suffix from lot of places.
- ([#1769](https://github.com/mesg-foundation/engine/pull/1769)) Update cosmos events.
- ([#1773](https://github.com/mesg-foundation/engine/pull/1773)) Use new event API in e2e tests.
- ([#1775](https://github.com/mesg-foundation/engine/pull/1775)) Simplify container.
- ([#1781](https://github.com/mesg-foundation/engine/pull/1781)) Switch back to default staking coins of cosmos.

#### Fixed

- ([#1746](https://github.com/mesg-foundation/engine/pull/1746)) Add "dive" validation to all repeated message in the proto files.
- ([#1762](https://github.com/mesg-foundation/engine/pull/1762)) Fix account sequence desynchronisation when error.

#### Dependencies

- ([#1752](https://github.com/mesg-foundation/engine/pull/1752)) Update to cosmos-sdk v0.38.2 and tendermint v0.33.2.

## [v0.20.0](https://github.com/mesg-foundation/engine/releases/tag/v0.20.0)

#### Added

- ([#1688](https://github.com/mesg-foundation/engine/pull/1688)) Handle withdraw from resources.
- ([#1694](https://github.com/mesg-foundation/engine/pull/1694)) Add build multi-platform cmds script to CI.
- ([#1699](https://github.com/mesg-foundation/engine/pull/1699)) Processes pay for executions.
- ([#1705](https://github.com/mesg-foundation/engine/pull/1705)) Add emitters to execution with consensus system.
- ([#1719](https://github.com/mesg-foundation/engine/pull/1719)) Add transaction's log to returned error.
- ([#1723](https://github.com/mesg-foundation/engine/pull/1723)) Add address to process.
- ([#1728](https://github.com/mesg-foundation/engine/pull/1728)) Add address to service, runner, execution, and execution.
- ([#1729](https://github.com/mesg-foundation/engine/pull/1729)) Wrap errors of modules' keeper, querier and msg to sent them to clients.

#### Dependencies

- ([#1668](https://github.com/mesg-foundation/engine/pull/1668)) Bump github.com/go-kit/kit from 0.9.0 to 0.10.0.
- ([#1674](https://github.com/mesg-foundation/engine/pull/1674)) Bump github.com/stretchr/testify from 1.4.0 to 1.5.1.
- ([#1677](https://github.com/mesg-foundation/engine/pull/1677)) Bump github.com/spf13/cobra from 0.0.5 to 0.0.6.
- ([#1686](https://github.com/mesg-foundation/engine/pull/1686)) Bump github.com/golang/protobuf from 1.3.3 to 1.3.4.
- ([#1697](https://github.com/mesg-foundation/engine/pull/1697)) Bump github.com/pseudomuto/protoc-gen-doc from 1.3.0 to 1.3.1.
- ([#1702](https://github.com/mesg-foundation/engine/pull/1702)) Bump github.com/prometheus/client_golang from 1.4.1 to 1.5.0.

## [v0.19.0](https://github.com/mesg-foundation/engine/releases/tag/v0.19.0)

#### Added

- ([#1654](https://github.com/mesg-foundation/engine/pull/1654)) Add e2e lcd tests for ownerships.
- ([#1663](https://github.com/mesg-foundation/engine/pull/1663)) Add e2e tests using the light client daemon of Cosmos.
- ([#1670](https://github.com/mesg-foundation/engine/pull/1670)) Introduce execution payment.
- ([#1680](https://github.com/mesg-foundation/engine/pull/1680)) Add block height to execution data.
- ([#1689](https://github.com/mesg-foundation/engine/pull/1689)) Add pagination system to http endpoint executions/list.
- ([#1691](https://github.com/mesg-foundation/engine/pull/1691)) Add min price for execution.

#### Changed

- ([#1640](https://github.com/mesg-foundation/engine/pull/1640)) Update to cosmos v0.38 and tendermint v0.33.
- ([#1645](https://github.com/mesg-foundation/engine/pull/1645)) Add trusted config to LCD server and index all events.
- ([#1650](https://github.com/mesg-foundation/engine/pull/1650)) Refactor cosmos modules.
- ([#1651](https://github.com/mesg-foundation/engine/pull/1651)) Refactor ownership module to cosmos style.
- ([#1653](https://github.com/mesg-foundation/engine/pull/1653)) Refactor instance module to cosmos style.
- ([#1657](https://github.com/mesg-foundation/engine/pull/1657)) Refactor runner module to cosmos style.
- ([#1658](https://github.com/mesg-foundation/engine/pull/1658)) Refactor process module to cosmos style.
- ([#1661](https://github.com/mesg-foundation/engine/pull/1661)) Refactor service module to cosmos style.
- ([#1662](https://github.com/mesg-foundation/engine/pull/1662)) Refactor execution module to cosmos style.
- ([#1669](https://github.com/mesg-foundation/engine/pull/1669)) Use disable-all instead of enable-all in golangci config.

#### Fixed

- ([#1682](https://github.com/mesg-foundation/engine/pull/1682)) Fix Tendermint graceful stop.
- ([#1683](https://github.com/mesg-foundation/engine/pull/1683)) Call end block function in each cosmos modules.

#### Dependencies

- ([#1646](https://github.com/mesg-foundation/engine/pull/1646)) Bump google.golang.org/grpc from 1.27.0 to 1.27.1.
- ([#1648](https://github.com/mesg-foundation/engine/pull/1648)) Bump github.com/prometheus/client_golang from 1.4.0 to 1.4.1.
- ([#1656](https://github.com/mesg-foundation/engine/pull/1656)) Bump github.com/cosmos/cosmos-sdk from 0.38.0 to 0.38.1.

## [v0.18.3](https://github.com/mesg-foundation/engine/releases/tag/v0.18.3)

#### Added

- ([#1630](https://github.com/mesg-foundation/engine/pull/1630)) Start cosmos lcd server with the engine.
- ([#1632](https://github.com/mesg-foundation/engine/pull/1632)) Add monitoring to dev engine.

#### Changed

- ([#1554](https://github.com/mesg-foundation/engine/pull/1554)) Cosmos implementation simplification.
- ([#1617](https://github.com/mesg-foundation/engine/pull/1617)) Use cosmos event manager like standard cosmos modules are doing.
- ([#1636](https://github.com/mesg-foundation/engine/pull/1636)) Reorganise monitoring configs and add default dashboards.
- ([#1641](https://github.com/mesg-foundation/engine/pull/1641)) Update metrics.

#### Fixed

- ([#1599](https://github.com/mesg-foundation/engine/pull/1599)) Move AccNumber and AccIndex to the config.
- ([#1625](https://github.com/mesg-foundation/engine/pull/1625)) Remove useless volume definition in the tool dockerfile.
- ([#1626](https://github.com/mesg-foundation/engine/pull/1626)) Emit only one event with multiple attributes in cosmos module handler.
- ([#1627](https://github.com/mesg-foundation/engine/pull/1627)) Simplify txbuilder and improve keybase signing.

#### Dependencies

- ([#1576](https://github.com/mesg-foundation/engine/pull/1576)) Bump gopkg.in/go-playground/validator.v9 from 9.30.2 to 9.31.0.
- ([#1610](https://github.com/mesg-foundation/engine/pull/1610)) Bump github.com/spf13/viper from 1.6.1 to 1.6.2.
- ([#1614](https://github.com/mesg-foundation/engine/pull/1614)) Bump gopkg.in/yaml.v2 from 2.2.7 to 2.2.8.
- ([#1616](https://github.com/mesg-foundation/engine/pull/1616)) Bump github.com/prometheus/client_golang from 1.1.0 to 1.4.0.
- ([#1618](https://github.com/mesg-foundation/engine/pull/1618)) Update to cosmos v0.37.6 and tendermint v0.32.9.
- ([#1619](https://github.com/mesg-foundation/engine/pull/1619)) Bump google.golang.org/grpc from 1.25.1 to 1.27.0.
- ([#1623](https://github.com/mesg-foundation/engine/pull/1623)) Bump github.com/golang/protobuf from 1.3.2 to 1.3.3.
- ([#1628](https://github.com/mesg-foundation/engine/pull/1628)) Bump github.com/grpc-ecosystem/go-grpc-middleware from 1.1.0 to 1.2.0.

## [v0.18.2](https://github.com/mesg-foundation/engine/releases/tag/v0.18.2)

#### Fixed

- ([#1605](https://github.com/mesg-foundation/engine/pull/1605)) Override default staking module config.

## [v0.18.1](https://github.com/mesg-foundation/engine/releases/tag/v0.18.1)

#### Added

- ([#1580](https://github.com/mesg-foundation/engine/pull/1580)) Add execution metrics.
- ([#1595](https://github.com/mesg-foundation/engine/pull/1595)) Add cosmos cli for low-level utility functionality.

#### Changed

- ([#1572](https://github.com/mesg-foundation/engine/pull/1572)) Customize cosmos address prefix.
- ([#1601](https://github.com/mesg-foundation/engine/pull/1601)) Change default block time to 5sec and some token config.
- ([#1602](https://github.com/mesg-foundation/engine/pull/1602)) Change the mesg coin type to the registered one (470).

#### Fixed

- ([#1588](https://github.com/mesg-foundation/engine/pull/1588)) Fix supply module init.
- ([#1598](https://github.com/mesg-foundation/engine/pull/1598)) Fix account number used for signing tx.
- ([#1553](https://github.com/mesg-foundation/engine/pull/1553)) Add fees and estimation of tx's gas.

#### Removed

- ([#1579](https://github.com/mesg-foundation/engine/pull/1579)) Remove useless database config.

## [v0.18.0](https://github.com/mesg-foundation/engine/releases/tag/v0.18.0)

#### Breaking Changes

- ([#1536](https://github.com/mesg-foundation/engine/pull/1536)) Update process structure: rename process's key to name and move node.*.key into node.
- ([#1540](https://github.com/mesg-foundation/engine/pull/1540)) Implement nested data map in process and change process map structure.
- ([#1545](https://github.com/mesg-foundation/engine/pull/1545)) Add path system to reference nested data in process.

#### Added

- ([#1527](https://github.com/mesg-foundation/engine/pull/1527)) Decentralization of the processes.
- ([#1519](https://github.com/mesg-foundation/engine/pull/1519)) Add data validation to all resources.
- ([#1567](https://github.com/mesg-foundation/engine/pull/1567)) Add validation on process node to check reference to non-task node.
- ([#1501](https://github.com/mesg-foundation/engine/pull/1501)) Add e2e tests on process and orchestrator. ([#1544](https://github.com/mesg-foundation/engine/pull/1544)). ([#1546](https://github.com/mesg-foundation/engine/pull/1546)). ([#1547](https://github.com/mesg-foundation/engine/pull/1547)). ([#1575](https://github.com/mesg-foundation/engine/pull/1575)).

#### Changed

- ([#1490](https://github.com/mesg-foundation/engine/pull/1490)) More verbose error in service backend.

#### Fixed

- ([#1532](https://github.com/mesg-foundation/engine/pull/1532)) Fix blocked logs by forcing to a maximum 10,000 lines in dev script.
- ([#1543](https://github.com/mesg-foundation/engine/pull/1543)) Fix concurrent transaction signing by adding a mutex to the keybase.
- ([#1556](https://github.com/mesg-foundation/engine/pull/1556)) Fix account sequence when signing multiple transactions.

#### Removed

- ([#1555](https://github.com/mesg-foundation/engine/pull/1555)) Remove account sdk and credential system.

## [v0.17.0](https://github.com/mesg-foundation/engine/releases/tag/v0.17.0)

#### Breaking Changes

- ([#1456](https://github.com/mesg-foundation/engine/pull/1456)) Update Instance SDK and create Runner SDK.
- ([#1459](https://github.com/mesg-foundation/engine/pull/1459)) Single cosmos account.
- ([#1463](https://github.com/mesg-foundation/engine/pull/1463)) Execution SDK on Cosmos.
- ([#1473](https://github.com/mesg-foundation/engine/pull/1473)) Use filter in list instance api.
- ([#1499](https://github.com/mesg-foundation/engine/pull/1499)) Improve use of proto Value and remove one of on UpdateExecutionRequest.
- ([#1513](https://github.com/mesg-foundation/engine/pull/1513)) Remove service env MESG_TOKEN.
- ([#1514](https://github.com/mesg-foundation/engine/pull/1514)) Remove service env MESG_TOKEN.

#### Added

- ([#1435](https://github.com/mesg-foundation/engine/pull/1435)) Prepare setup for running e2e tests.
- ([#1439](https://github.com/mesg-foundation/engine/pull/1439)) Setup up account test in E2E.
- ([#1448](https://github.com/mesg-foundation/engine/pull/1448)) Inject env MESG_INSTANCE_HASH in services.
- ([#1449](https://github.com/mesg-foundation/engine/pull/1449)) Add config log format validation.
- ([#1450](https://github.com/mesg-foundation/engine/pull/1450)) Add more explicit error on service create.
- ([#1451](https://github.com/mesg-foundation/engine/pull/1451)) Validate service nested object params.
- ([#1489](https://github.com/mesg-foundation/engine/pull/1489)) Add account mnemonic in config.
- ([#1500](https://github.com/mesg-foundation/engine/pull/1500)) Add e2e tests of a complex service.
- ([#1502](https://github.com/mesg-foundation/engine/pull/1502)) Add amino encoding functionality to proto Struct.
- ([#1503](https://github.com/mesg-foundation/engine/pull/1503)) Add small test to make sure hashstruct is sorting the map and struct keys.
- ([#1512](https://github.com/mesg-foundation/engine/pull/1512)) Inject runner hash in docker services.

#### Changed

- ([#1432](https://github.com/mesg-foundation/engine/pull/1432)) Bump gopkg.in/yaml.v2 from 2.2.2 to 2.2.4.
- ([#1437](https://github.com/mesg-foundation/engine/pull/1437)) Remove dep from lint and add dep in docker-dev in makefile.
- ([#1452](https://github.com/mesg-foundation/engine/pull/1452)) Add EnvHash to instance and calculate the hash based on it.
- ([#1461](https://github.com/mesg-foundation/engine/pull/1461)) Refactor instance sdk containers function.
- ([#1474](https://github.com/mesg-foundation/engine/pull/1474)) Update golangci-lint to 1.21.0.
- ([#1485](https://github.com/mesg-foundation/engine/pull/1485)) Increase gas limit tx.
- ([#1488](https://github.com/mesg-foundation/engine/pull/1488)) Refactor database.
- ([#1491](https://github.com/mesg-foundation/engine/pull/1491)) Filter instances on backend side.

#### Fixed

- ([#1462](https://github.com/mesg-foundation/engine/pull/1462)) Use sequence number from cosmos auth module.
- ([#1479](https://github.com/mesg-foundation/engine/pull/1479)) Return proper error code on e2e fail.
- ([#1480](https://github.com/mesg-foundation/engine/pull/1480)) Use a single cosmos codec.
- ([#1481](https://github.com/mesg-foundation/engine/pull/1481)) export MESG_PATH for e2e tests.
- ([#1483](https://github.com/mesg-foundation/engine/pull/1483)) Disable go test cache for e2e test.
- ([#1486](https://github.com/mesg-foundation/engine/pull/1486)) Fix 2 error messages.
- ([#1504](https://github.com/mesg-foundation/engine/pull/1504)) Fix hash length.
- ([#1510](https://github.com/mesg-foundation/engine/pull/1510)) Return not found error in backend get functions when resource doesn't exist.
- ([#1511](https://github.com/mesg-foundation/engine/pull/1511)) Fix bug when Value_NullValue is not taking into account.
- ([#1521](https://github.com/mesg-foundation/engine/pull/1521)) Fix filter in runner sdk and check if runner exists.
- ([#1522](https://github.com/mesg-foundation/engine/pull/1522)) Fix typos.

#### Removed

- ([#1468](https://github.com/mesg-foundation/engine/pull/1468)) Remove public error AlreadyExistsError from SDKs.
- ([#1469](https://github.com/mesg-foundation/engine/pull/1469)) Remove servicesdk interface.
- ([#1482](https://github.com/mesg-foundation/engine/pull/1482)) Remove runner db to use directly cosmos kvstore in runner sdk.
- ([#1508](https://github.com/mesg-foundation/engine/pull/1508)) Remove deprecated database files.
- ([#1516](https://github.com/mesg-foundation/engine/pull/1516)) Remove database/store package.

## [v0.16.0](https://github.com/mesg-foundation/engine/releases/tag/v0.16.0)

#### Breaking Changes

- ([#1361](https://github.com/mesg-foundation/engine/pull/1361)) Remove system services and deprecated service sdk.
- ([#1364](https://github.com/mesg-foundation/engine/pull/1364)) Remove core api.
- ([#1375](https://github.com/mesg-foundation/engine/pull/1375)) Simplify service hash calculation.
- ([#1381](https://github.com/mesg-foundation/engine/pull/1381)) Remove service deletion function.
- ([#1416](https://github.com/mesg-foundation/engine/pull/1416)) Read config from file instead of env.

#### Added

- ([#1355](https://github.com/mesg-foundation/engine/pull/1355)) Add credential using grpc metadata.
- ([#1368](https://github.com/mesg-foundation/engine/pull/1368)) Add ownership SDK and API.
- ([#1380](https://github.com/mesg-foundation/engine/pull/1380)) Add configuration validation.
- ([#1384](https://github.com/mesg-foundation/engine/pull/1384)) Add service exists api.

#### Changed

- ([#1357](https://github.com/mesg-foundation/engine/pull/1357)) Update dependencies.
- ([#1371](https://github.com/mesg-foundation/engine/pull/1371)) Remove shared private info of the genesis account to generate the genesis.
- ([#1391](https://github.com/mesg-foundation/engine/pull/1391)) Improve generate genesis script.
- ([#1400](https://github.com/mesg-foundation/engine/pull/1400)) Container package refactor.
- ([#1402](https://github.com/mesg-foundation/engine/pull/1402)) Run scripts from Makefile.
- ([#1412](https://github.com/mesg-foundation/engine/pull/1412)) Load genesis from file and generate it if it doesn't exist.
- ([#1430](https://github.com/mesg-foundation/engine/pull/1430)) Set block creation time to 1s instead of 10s.
- ([#1431](https://github.com/mesg-foundation/engine/pull/1431)) Update dependencies.

#### Fixed

- ([#1360](https://github.com/mesg-foundation/engine/pull/1360)) Fix copy of folder in dev script.

#### Removed

- ([#1370](https://github.com/mesg-foundation/engine/pull/1370)) Remove system service sources.
- ([#1373](https://github.com/mesg-foundation/engine/pull/1373)) Remove not used service-test.
- ([#1414](https://github.com/mesg-foundation/engine/pull/1414)) Remove not used functions from xos package.
- ([#1428](https://github.com/mesg-foundation/engine/pull/1428)) Remove deprecated `MESG_ENDPOINT_TCP` from service env.

## [v0.15.0](https://github.com/mesg-foundation/engine/releases/tag/v0.15.0)

#### Breaking Changes

- ([#1329](https://github.com/mesg-foundation/engine/pull/1329)) Calculate process hash based on the nodes.
- ([#1335](https://github.com/mesg-foundation/engine/pull/1335)) Remove plural form from instance list proto service.
- ([#1337](https://github.com/mesg-foundation/engine/pull/1337)) Namespace protobuf types and api.

#### Experimental

- ([#1267](https://github.com/mesg-foundation/engine/pull/1267)) Implement service SDK with Cosmos.
- ([#1308](https://github.com/mesg-foundation/engine/pull/1308)) Simplify cosmos keybase by just creating a NewMnemonic helper.
- ([#1331](https://github.com/mesg-foundation/engine/pull/1331)) Add --exp to dev command.
- ([#1343](https://github.com/mesg-foundation/engine/pull/1343)) Add account sdk and its api.

#### Added

- ([#1316](https://github.com/mesg-foundation/engine/pull/1316)) Add constant to process map.
- ([#1317](https://github.com/mesg-foundation/engine/pull/1317)) Add tests for orchestrator.
- ([#1342](https://github.com/mesg-foundation/engine/pull/1342)) Build and publish docker image with minor-only tag.

#### Changed

- ([#1305](https://github.com/mesg-foundation/engine/pull/1305)) Use protobuf struct in create function of service SDK.
- ([#1321](https://github.com/mesg-foundation/engine/pull/1321)) Only force hash calculation on interface containing struct.
- ([#1322](https://github.com/mesg-foundation/engine/pull/1322)) Update golang to 1.13.
- ([#1323](https://github.com/mesg-foundation/engine/pull/1323)) Remove unused docker api mocks.
- ([#1347](https://github.com/mesg-foundation/engine/pull/1347)) Update marketplace system service with latest mesg-js.

## [v0.14.2](https://github.com/mesg-foundation/engine/releases/tag/v0.14.2)

#### Fixed

- ([#1339](https://github.com/mesg-foundation/engine/pull/1339)) Fix hash.Unmarshal when using proto.Unmarshal.

## [v0.14.1](https://github.com/mesg-foundation/engine/releases/tag/v0.14.1)

#### Fixed

- ([#1318](https://github.com/mesg-foundation/engine/pull/1318)) Fix hash calculation based on proto.Struct.
- ([#1320](https://github.com/mesg-foundation/engine/pull/1320)) Use proto.Struct in the orchestrator. ([#1327](https://github.com/mesg-foundation/engine/pull/1327)).

#### Changed

- ([#1309](https://github.com/mesg-foundation/engine/pull/1309)) Do not convert proto.Struct to map in order to Validate it.

#### Removed

## [v0.14.0](https://github.com/mesg-foundation/engine/releases/tag/v0.14.0)

#### Breaking Changes

- ([#1258](https://github.com/mesg-foundation/engine/pull/1258)) Pure graph implementation for the orchestration system.
- ([#1278](https://github.com/mesg-foundation/engine/pull/1278)) Rename workflow to process.
- ([#1290](https://github.com/mesg-foundation/engine/pull/1290)) Use gogo protobuf.
- ([#1292](https://github.com/mesg-foundation/engine/pull/1292)) Set nullabe to false for service configuration.
- ([#1301](https://github.com/mesg-foundation/engine/pull/1301)) Update database path to ensure no issue on startup.

#### Added

- ([#1241](https://github.com/mesg-foundation/engine/pull/1241)) Process input resolution.

#### Changed

- ([#1254](https://github.com/mesg-foundation/engine/pull/1254)) Remove database ErrNotFound. Introduce Database Exist function.
- ([#1260](https://github.com/mesg-foundation/engine/pull/1260)) Keep only one Service struct (the one generated from protobuf definition).
- ([#1264](https://github.com/mesg-foundation/engine/pull/1264)) Refactor/remove custom service type no custom type hash.
- ([#1269](https://github.com/mesg-foundation/engine/pull/1269)) Change service.proto package back to types.
- ([#1270](https://github.com/mesg-foundation/engine/pull/1270)) Change proto hash type from string to bytes.
- ([#1285](https://github.com/mesg-foundation/engine/pull/1285)) Bump github.com/gogo/protobuf from 1.2.1 to 1.3.0.
- ([#1286](https://github.com/mesg-foundation/engine/pull/1286)) Fix marshal issue with process filter.
- ([#1287](https://github.com/mesg-foundation/engine/pull/1287)) Update system services to work with latest gRPC API.
- ([#1288](https://github.com/mesg-foundation/engine/pull/1288)) Fix that hash of info API were still in string instead of bytes.
- ([#1291](https://github.com/mesg-foundation/engine/pull/1291)) Customize protobuf messages .
- ([#1293](https://github.com/mesg-foundation/engine/pull/1293)) Replace instance struct with protobuf.
- ([#1295](https://github.com/mesg-foundation/engine/pull/1295)) Remove old generated instance struct.
- ([#1296](https://github.com/mesg-foundation/engine/pull/1296)) Remove and refactor rest protobuf types.
- ([#1298](https://github.com/mesg-foundation/engine/pull/1298)) Replace cnf structhash with custom one.
- ([#1310](https://github.com/mesg-foundation/engine/pull/1310)) Update gogo protobuf with fixed json and custom tags generation.
- ([#1311](https://github.com/mesg-foundation/engine/pull/1311)) Use protobuf marshal/unmarshal in process database.
- ([#1312](https://github.com/mesg-foundation/engine/pull/1312)) Use protobuf marshal/unmarshal in execution database.

#### Experimental

- ([#1243](https://github.com/mesg-foundation/engine/pull/1243)) Add store package.
- ([#1244](https://github.com/mesg-foundation/engine/pull/1244)) Update service db to use store.
- ([#1245](https://github.com/mesg-foundation/engine/pull/1245)) Bump github.com/cosmos/cosmos-sdk from 0.36.0 to 0.37.0.
- ([#1248](https://github.com/mesg-foundation/engine/pull/1248)) Add Cosmos helpers.
- ([#1249](https://github.com/mesg-foundation/engine/pull/1249)) SDK loads the App modules.
- ([#1250](https://github.com/mesg-foundation/engine/pull/1250)) Transform service sdk to accept cosmos service step 1.
- ([#1252](https://github.com/mesg-foundation/engine/pull/1252)) Add service sdk cosmos version step 2.
- ([#1255](https://github.com/mesg-foundation/engine/pull/1255)) Create tendermint client.
- ([#1257](https://github.com/mesg-foundation/engine/pull/1257)) Recover from panic in cosmos store.
- ([#1268](https://github.com/mesg-foundation/engine/pull/1268)) Make cosmos store iterator works the same way as goleveldb iterator.
- ([#1271](https://github.com/mesg-foundation/engine/pull/1271)) Bump github.com/tendermint/tendermint from 0.32.2 to 0.32.3.
- ([#1284](https://github.com/mesg-foundation/engine/pull/1284)) Add test for both cosmos store and leveldb store.
- ([#1297](https://github.com/mesg-foundation/engine/pull/1297)) Add explicit marshal and unmarshal function to service db.
- ([#1299](https://github.com/mesg-foundation/engine/pull/1299)) Create a cosmos app factory struct.
- ([#1300](https://github.com/mesg-foundation/engine/pull/1300)) Prepare cosmos client.

## [v0.13.0](https://github.com/mesg-foundation/engine/releases/tag/v0.13.0)

#### Added

- ([#1201](https://github.com/mesg-foundation/engine/pull/1201)) Multi step workflow.
- ([#1213](https://github.com/mesg-foundation/engine/pull/1213)) Add module name to logs.
- ([#1223](https://github.com/mesg-foundation/engine/pull/1223)) Workflow validation.

#### Changed

- ([#1204](https://github.com/mesg-foundation/engine/pull/1204)) Move workflow to its own database/api.
- ([#1208](https://github.com/mesg-foundation/engine/pull/1208)) Rename engine package to scheduler.
- ([#1209](https://github.com/mesg-foundation/engine/pull/1209)) Workflow graph structure.
- ([#1214](https://github.com/mesg-foundation/engine/pull/1214)) Create struct service.Configuration.
- ([#1229](https://github.com/mesg-foundation/engine/pull/1229)) Remove trigger type for workflows.
- ([#1231](https://github.com/mesg-foundation/engine/pull/1231)) Use required_without validation instead of manual one.

#### Fixed

- ([#1207](https://github.com/mesg-foundation/engine/pull/1207)) Clean dev script.
- ([#1210](https://github.com/mesg-foundation/engine/pull/1210)) Improve regex used in dev script.
- ([#1215](https://github.com/mesg-foundation/engine/pull/1215)) Remove unused variable and improve error handling of config.setupServices function.
- ([#1217](https://github.com/mesg-foundation/engine/pull/1217)) Fix dev script.
- ([#1220](https://github.com/mesg-foundation/engine/pull/1220)) Remove regexp network filter from dev script.

#### Removed

- ([#1230](https://github.com/mesg-foundation/engine/pull/1230)) Remove git related package.

#### Experimental

- ([#1211](https://github.com/mesg-foundation/engine/pull/1211)) Cosmos integration.
- ([#1218](https://github.com/mesg-foundation/engine/pull/1218)) Set tm config in config package and pass it on node creation.
- ([#1219](https://github.com/mesg-foundation/engine/pull/1219)) Add cosmos modules.
- ([#1226](https://github.com/mesg-foundation/engine/pull/1226)) update Cosmos to v0.36.
- ([#1227](https://github.com/mesg-foundation/engine/pull/1227)) Possibility to enable/disable network when running the engine.

## [v0.12.1](https://github.com/mesg-foundation/engine/releases/tag/v0.12.1)

#### Fixed

- ([#1199](https://github.com/mesg-foundation/engine/pull/1199)) Fix issue with the `marketplace#getService` task.

## [v0.12.0](https://github.com/mesg-foundation/engine/releases/tag/v0.12.0)

#### Breaking Changes

- ([#1173](https://github.com/mesg-foundation/engine/pull/1173)) `mesg.hash` has been removed and replaced by `mesg.service` and `mesg.instance` in the docker labels.
- ([#1192](https://github.com/mesg-foundation/engine/pull/1192)) Replace JSON format with protobuf Struct on gRPC APIs for `inputs`, `outputs` and `data` parameters. ([#1198](https://github.com/mesg-foundation/engine/pull/1198)).

#### Features

- ([#1191](https://github.com/mesg-foundation/engine/pull/1191)) Put executions as failed when a service emits the wrong outputs.

#### Experimental

 - ([#1172](https://github.com/mesg-foundation/engine/pull/1172)) Implementation of workflow engine foundation. ([#1175](https://github.com/mesg-foundation/engine/pull/1175)). ([#1184](https://github.com/mesg-foundation/engine/pull/1184)). ([#1181](https://github.com/mesg-foundation/engine/pull/1181)). ([#1182](https://github.com/mesg-foundation/engine/pull/1182)). ([#1185](https://github.com/mesg-foundation/engine/pull/1185)). ([#1188](https://github.com/mesg-foundation/engine/pull/1188)). ([#1190](https://github.com/mesg-foundation/engine/pull/1190)).


## [v0.11.0](https://github.com/mesg-foundation/engine/releases/tag/v0.11.0)

### [Click here to see the release notes](https://forum.mesg.com/t/release-notes-of-engine-v0-11-cli-v1-1-and-js-library-v4/339)

#### Breaking Changes

- ([#1083](https://github.com/mesg-foundation/engine/pull/1083)) Remove old gRPC APIs.
- ([#1170](https://github.com/mesg-foundation/engine/pull/1170)) Change Service and Execution database version. You'll need to redeploy your services.

#### Added

- ([#1023](https://github.com/mesg-foundation/engine/pull/1023)) New gRPC Service APIs and SDK. ([#1066](https://github.com/mesg-foundation/engine/pull/1066)). ([#1067](https://github.com/mesg-foundation/engine/pull/1067)). ([#1068](https://github.com/mesg-foundation/engine/pull/1068)). ([#1071](https://github.com/mesg-foundation/engine/pull/1071)). ([#1077](https://github.com/mesg-foundation/engine/pull/1077)). ([#1097](https://github.com/mesg-foundation/engine/pull/1097)). ([#1099](https://github.com/mesg-foundation/engine/pull/1099)). ([#1105](https://github.com/mesg-foundation/engine/pull/1105)). ([#1107](https://github.com/mesg-foundation/engine/pull/1107)). ([#1110](https://github.com/mesg-foundation/engine/pull/1110)). ([#1112](https://github.com/mesg-foundation/engine/pull/1112)). ([#1113](https://github.com/mesg-foundation/engine/pull/1113)). ([#1128](https://github.com/mesg-foundation/engine/pull/1128)). ([#1153](https://github.com/mesg-foundation/engine/pull/1153)).
- ([#1033](https://github.com/mesg-foundation/engine/pull/1033)) New gRPC Instance APIs and SDK. ([#1034](https://github.com/mesg-foundation/engine/pull/1034)). ([#1035](https://github.com/mesg-foundation/engine/pull/1035)). ([#1036](https://github.com/mesg-foundation/engine/pull/1036)). ([#1037](https://github.com/mesg-foundation/engine/pull/1037)). ([#1060](https://github.com/mesg-foundation/engine/pull/1060)). ([#1074](https://github.com/mesg-foundation/engine/pull/1074)). ([#1075](https://github.com/mesg-foundation/engine/pull/1075)). ([#1076](https://github.com/mesg-foundation/engine/pull/1076)). ([#1078](https://github.com/mesg-foundation/engine/pull/1078)). ([#1079](https://github.com/mesg-foundation/engine/pull/1079)). ([#1087](https://github.com/mesg-foundation/engine/pull/1087)). ([#1102](https://github.com/mesg-foundation/engine/pull/1102)). ([#1106](https://github.com/mesg-foundation/engine/pull/1106)). ([#1109](https://github.com/mesg-foundation/engine/pull/1109)). ([#1117](https://github.com/mesg-foundation/engine/pull/1117)). ([#1122](https://github.com/mesg-foundation/engine/pull/1122)). ([#1137](https://github.com/mesg-foundation/engine/pull/1137)). ([#1138](https://github.com/mesg-foundation/engine/pull/1138)). ([#1156](https://github.com/mesg-foundation/engine/pull/1156)).
- ([#1043](https://github.com/mesg-foundation/engine/pull/1043)) New gRPC Execution APIs and SDK. ([#1052](https://github.com/mesg-foundation/engine/pull/1052)). ([#1064](https://github.com/mesg-foundation/engine/pull/1064)). ([#1065](https://github.com/mesg-foundation/engine/pull/1065)). ([#1070](https://github.com/mesg-foundation/engine/pull/1070)). ([#1124](https://github.com/mesg-foundation/engine/pull/1124)). ([#1132](https://github.com/mesg-foundation/engine/pull/1132)). ([#1135](https://github.com/mesg-foundation/engine/pull/1135)).
- ([#1053](https://github.com/mesg-foundation/engine/pull/1053)) New gRPC Event APIs and SDK. ([#1054](https://github.com/mesg-foundation/engine/pull/1054)). ([#1073](https://github.com/mesg-foundation/engine/pull/1073)). ([#1126](https://github.com/mesg-foundation/engine/pull/1126)). ([#1144](https://github.com/mesg-foundation/engine/pull/1144)). ([#1141](https://github.com/mesg-foundation/engine/pull/1141)).

#### Changed

- ([#1082](https://github.com/mesg-foundation/engine/pull/1082)) Server and SDK package cleanup. ([#1085](https://github.com/mesg-foundation/engine/pull/1085)). ([#1096](https://github.com/mesg-foundation/engine/pull/1096)).
- ([#1087](https://github.com/mesg-foundation/engine/pull/1087)) Update system service deployment system. ([#1156](https://github.com/mesg-foundation/engine/pull/1156)). ([#1160](https://github.com/mesg-foundation/engine/pull/1160)).
- ([#1094](https://github.com/mesg-foundation/engine/pull/1094)) Introduction of a Hash type to manage all encoding and decoding of hashes. ([#1072](https://github.com/mesg-foundation/engine/pull/1072)). ([#1098](https://github.com/mesg-foundation/engine/pull/1098)).
- ([#1045](https://github.com/mesg-foundation/engine/pull/1045)) Update system service Marketplace to use new gRPC APIs. ([#1166](https://github.com/mesg-foundation/engine/pull/1166)).
- ([#1148](https://github.com/mesg-foundation/engine/pull/1148)) Update system service EthWallet to use new gRPC APIs. ([#1167](https://github.com/mesg-foundation/engine/pull/1167)).
- ([#1151](https://github.com/mesg-foundation/engine/pull/1151)) Namespace simplification in package Container.

## [v0.10.1](https://github.com/mesg-foundation/engine/releases/tag/v0.10.1)

### [Click here to see the release notes](https://forum.mesg.com/t/mesg-engine-v0-10-js-cli-and-js-library-v3-0-0-release-notes/317)

#### Fixed

- ([#1050](https://github.com/mesg-foundation/engine/pull/1050)) Fix service Marketplace backward compatibility.

## [v0.10.0](https://github.com/mesg-foundation/engine/releases/tag/v0.10.0)

### [Click here to see the release notes](https://forum.mesg.com/t/mesg-engine-v0-10-js-cli-and-js-library-v3-0-0-release-notes/317)

#### Breaking Changes

- ([#928](https://github.com/mesg-foundation/engine/pull/928)) Rename `volumesfrom` to `volumesFrom` in services' `mesg.yml`.
- ([#974](https://github.com/mesg-foundation/engine/pull/974)) Reduce the number of service outputs to one. An output can still contain multiple parameters. ([#963](https://github.com/mesg-foundation/engine/pull/963)). ([#964](https://github.com/mesg-foundation/engine/pull/964)). ([#965](https://github.com/mesg-foundation/engine/pull/965)). ([#967](https://github.com/mesg-foundation/engine/pull/967)). ([#971](https://github.com/mesg-foundation/engine/pull/971)). ([#972](https://github.com/mesg-foundation/engine/pull/972)). ([#973](https://github.com/mesg-foundation/engine/pull/973)). ([#975](https://github.com/mesg-foundation/engine/pull/975)). ([#976](https://github.com/mesg-foundation/engine/pull/976)). ([#977](https://github.com/mesg-foundation/engine/pull/977)). ([#978](https://github.com/mesg-foundation/engine/pull/978)). ([#979](https://github.com/mesg-foundation/engine/pull/979)). ([#980](https://github.com/mesg-foundation/engine/pull/980)). ([#981](https://github.com/mesg-foundation/engine/pull/981)). ([#982](https://github.com/mesg-foundation/engine/pull/982)). ([#983](https://github.com/mesg-foundation/engine/pull/983)). ([#984](https://github.com/mesg-foundation/engine/pull/984)). ([#987](https://github.com/mesg-foundation/engine/pull/987)). ([#988](https://github.com/mesg-foundation/engine/pull/988)). ([#1028](https://github.com/mesg-foundation/engine/pull/1028)).
- ([#791](https://github.com/mesg-foundation/engine/pull/791)) Remove CLI from the repository. The new cli is available on a [dedicated repository](https://github.com/mesg-foundation/cli). ([#995](https://github.com/mesg-foundation/engine/pull/995)). ([#996](https://github.com/mesg-foundation/engine/pull/996)).
- ([#991](https://github.com/mesg-foundation/engine/pull/991)) Rename Core to Engine. ([#968](https://github.com/mesg-foundation/engine/pull/968)). ([#970](https://github.com/mesg-foundation/engine/pull/970)). ([#1002](https://github.com/mesg-foundation/engine/pull/1002)). ([#1003](https://github.com/mesg-foundation/engine/pull/1003)). ([#1004](https://github.com/mesg-foundation/engine/pull/1004)). ([#1020](https://github.com/mesg-foundation/engine/pull/1020)).
- ([#1032](https://github.com/mesg-foundation/engine/pull/1032)) Simplify Engine configs. ([#1038](https://github.com/mesg-foundation/engine/pull/1038)).

#### Added

- ([#1014](https://github.com/mesg-foundation/engine/pull/1014)) Introduce new deployment api. Only available for development purpose.

#### Changed

- ([#994](https://github.com/mesg-foundation/engine/pull/994)) Update execution database and api. ([#1006](https://github.com/mesg-foundation/engine/pull/1006)). ([#1007](https://github.com/mesg-foundation/engine/pull/1007)). ([#1041](https://github.com/mesg-foundation/engine/pull/1041)).
- ([#997](https://github.com/mesg-foundation/engine/pull/997)) Rename package `api` to `sdk`.
- ([#998](https://github.com/mesg-foundation/engine/pull/998)) Rename package `interface` to `server`.

#### Fixed

- ([#955](https://github.com/mesg-foundation/engine/pull/955)) Catch error when the volume is deleted and it did not exist.

#### Removed

- [#1001](https://github.com/mesg-foundation/engine/pull/1001) Remove web3 http request logs from Marketplace.

## [v0.9.1](https://github.com/mesg-foundation/engine/releases/tag/v0.9.1)

### [Click here to see the release notes](https://forum.mesg.com/t/mesg-core-v0-9-1-release-notes/293)

#### Added

- ([#993](https://github.com/mesg-foundation/engine/pull/993)) Display web3 http request logs in Marketplace.

#### Changed

- ([#949](https://github.com/mesg-foundation/engine/pull/949)) Use MESG's IPFS node in CLI.

#### Fixed

- ([#930](https://github.com/mesg-foundation/engine/pull/930)) Improve generated README when using command `service gen-doc`. ([#948](https://github.com/mesg-foundation/engine/pull/948)). ([#960](https://github.com/mesg-foundation/engine/pull/960)).
- ([#934](https://github.com/mesg-foundation/engine/pull/934)) Return error when an image is passed on a mesg.yml definition.
- ([#929](https://github.com/mesg-foundation/engine/pull/929)) Show more verbose error when deploying service.

#### Documentation

- ([#953](https://github.com/mesg-foundation/engine/pull/953)) Fix links to docs

## [v0.9.0](https://github.com/mesg-foundation/engine/releases/tag/v0.9.0)

### [Click here to see the release notes](https://forum.mesg.com/t/mesg-core-v0-9-release-notes/273)

#### Breaking Changes

- ([#731](https://github.com/mesg-foundation/engine/pull/731)) Deterministic service hash. ([#804](https://github.com/mesg-foundation/engine/pull/804)). ([#877](https://github.com/mesg-foundation/engine/pull/877)).
- ([#801](https://github.com/mesg-foundation/engine/pull/801)) Add Service's hash to reply of Deploy API
- ([#849](https://github.com/mesg-foundation/engine/pull/849)) Use base58 to encode service hash.
- ([#860](https://github.com/mesg-foundation/engine/pull/860)) Separate service's configuration from service's dependencies. ([#880](https://github.com/mesg-foundation/engine/pull/880)).
- ([#866](https://github.com/mesg-foundation/engine/pull/866)) Rename service's `volumesfrom` property.
- ([#905](https://github.com/mesg-foundation/engine/pull/905)) Add version to database path to prevent decoding error and loss of data.

#### Added

- ([#535](https://github.com/mesg-foundation/engine/pull/535)) Run MESG with MESG Services! [#567](https://github.com/mesg-foundation/engine/pull/567).
- ([#755](https://github.com/mesg-foundation/engine/pull/755)) Implementation of the MESG Marketplace. The Marketplace allows the distribution and reutilization of MESG Services. Check `mesg-core marketplace` commands. ([#778](https://github.com/mesg-foundation/engine/pull/778)). ([#788](https://github.com/mesg-foundation/engine/pull/788)). ([#810](https://github.com/mesg-foundation/engine/pull/810)). ([#817](https://github.com/mesg-foundation/engine/pull/817)). ([#826](https://github.com/mesg-foundation/engine/pull/826)). ([#828](https://github.com/mesg-foundation/engine/pull/828)). ([#829](https://github.com/mesg-foundation/engine/pull/829)). ([#837](https://github.com/mesg-foundation/engine/pull/837)). ([#843](https://github.com/mesg-foundation/engine/pull/843)). ([#844](https://github.com/mesg-foundation/engine/pull/844)). ([#845](https://github.com/mesg-foundation/engine/pull/845)). ([#853](https://github.com/mesg-foundation/engine/pull/853)). ([#854](https://github.com/mesg-foundation/engine/pull/854)). ([#863](https://github.com/mesg-foundation/engine/pull/863)). ([#864](https://github.com/mesg-foundation/engine/pull/864)). ([#868](https://github.com/mesg-foundation/engine/pull/868)). ([#874](https://github.com/mesg-foundation/engine/pull/874)). ([#883](https://github.com/mesg-foundation/engine/pull/883)). ([#899](https://github.com/mesg-foundation/engine/pull/899)). ([#898](https://github.com/mesg-foundation/engine/pull/898)). ([#897](https://github.com/mesg-foundation/engine/pull/897)). ([#896](https://github.com/mesg-foundation/engine/pull/896)). ([#902](https://github.com/mesg-foundation/engine/pull/902)). ([#901](https://github.com/mesg-foundation/engine/pull/901)). ([#906](https://github.com/mesg-foundation/engine/pull/906)). ([#907](https://github.com/mesg-foundation/engine/pull/907)). ([#908](https://github.com/mesg-foundation/engine/pull/908)). ([#909](https://github.com/mesg-foundation/engine/pull/909)). ([#924](https://github.com/mesg-foundation/engine/pull/924)). ([#926](https://github.com/mesg-foundation/engine/pull/926)). ([#927](https://github.com/mesg-foundation/engine/pull/927)). ([#936](https://github.com/mesg-foundation/engine/pull/936)). ([#938](https://github.com/mesg-foundation/engine/pull/938)). ([#939](https://github.com/mesg-foundation/engine/pull/939)). ([#942](https://github.com/mesg-foundation/engine/pull/942)). ([#943](https://github.com/mesg-foundation/engine/pull/943)).
- ([#757](https://github.com/mesg-foundation/engine/pull/757)) Read `.dockerignore` in dev and deploy commands.
- ([#779](https://github.com/mesg-foundation/engine/pull/779)) Implementation of the MESG Wallet. Check `mesg-core wallet`. ([#807](https://github.com/mesg-foundation/engine/pull/807)). ([#809](https://github.com/mesg-foundation/engine/pull/809)). ([#812](https://github.com/mesg-foundation/engine/pull/812)). ([#937](https://github.com/mesg-foundation/engine/pull/937)).
- ([#781](https://github.com/mesg-foundation/engine/pull/781)) Improve validation of service definition. ([#869](https://github.com/mesg-foundation/engine/pull/869)).

#### Changed

- ([#823](https://github.com/mesg-foundation/engine/pull/823)) Improve commands `service init` and `service gendoc`.
- ([#875](https://github.com/mesg-foundation/engine/pull/875)) Improve JSON parsing error message.
- ([#790](https://github.com/mesg-foundation/engine/pull/790)) Refactor. ([#792](https://github.com/mesg-foundation/engine/pull/792)). ([#816](https://github.com/mesg-foundation/engine/pull/816)). ([#805](https://github.com/mesg-foundation/engine/pull/805)). ([#813](https://github.com/mesg-foundation/engine/pull/813)). ([#839](https://github.com/mesg-foundation/engine/pull/839)). ([#847](https://github.com/mesg-foundation/engine/pull/847)). ([#850](https://github.com/mesg-foundation/engine/pull/850)). ([#852](https://github.com/mesg-foundation/engine/pull/852)). ([#855](https://github.com/mesg-foundation/engine/pull/855)). ([#858](https://github.com/mesg-foundation/engine/pull/858)). ([#867](https://github.com/mesg-foundation/engine/pull/867)). ([#859](https://github.com/mesg-foundation/engine/pull/859)). ([#870](https://github.com/mesg-foundation/engine/pull/870)). ([#871](https://github.com/mesg-foundation/engine/pull/871)). ([#872](https://github.com/mesg-foundation/engine/pull/872)). ([#873](https://github.com/mesg-foundation/engine/pull/873)). ([#881](https://github.com/mesg-foundation/engine/pull/881)). ([#893](https://github.com/mesg-foundation/engine/pull/893)). ([#892](https://github.com/mesg-foundation/engine/pull/892)). ([#891](https://github.com/mesg-foundation/engine/pull/891)). ([#890](https://github.com/mesg-foundation/engine/pull/890)). ([#889](https://github.com/mesg-foundation/engine/pull/889)). ([#888](https://github.com/mesg-foundation/engine/pull/888)). ([#886](https://github.com/mesg-foundation/engine/pull/886)). ([#885](https://github.com/mesg-foundation/engine/pull/885)). ([#884](https://github.com/mesg-foundation/engine/pull/884)). ([#903](https://github.com/mesg-foundation/engine/pull/903)). ([#919](https://github.com/mesg-foundation/engine/pull/919)).

#### Fixed

- ([#771](https://github.com/mesg-foundation/engine/pull/771)) Fix gRPC stream acknowledgement.
- ([#772](https://github.com/mesg-foundation/engine/pull/772)) Improve command logs errors.
- ([#820](https://github.com/mesg-foundation/engine/pull/820)) Fix container package.

## [v0.8.1](https://github.com/mesg-foundation/engine/releases/tag/v0.8.1)

### [Click here to see the release notes](https://forum.mesg.com/t/mesg-core-v0-8-1-release-notes/249)

#### Fixed

- ([#774](https://github.com/mesg-foundation/engine/pull/774)) Update keep alive of client to 5min to prevent spamming the server.

#### Documentation

- ([#762](https://github.com/mesg-foundation/engine/pull/762)) Fix and improve guide start. ([#763](https://github.com/mesg-foundation/engine/pull/763)).

## [v0.8.0](https://github.com/mesg-foundation/engine/releases/tag/v0.8.0)

### [Click here to see the release notes](https://forum.mesg.com/t/mesg-core-v0-8-release-notes/239/2)

#### Added

- ([#690](https://github.com/mesg-foundation/engine/pull/690)) Support service deployments from tarball urls.
- ([#732](https://github.com/mesg-foundation/engine/pull/732)) Support multiple service id or hash for commands `service start` and `service stop`.
- ([#726](https://github.com/mesg-foundation/engine/pull/726)) Add flag to command `start` to force colors in logs of Core.

#### Changed

- ([#734](https://github.com/mesg-foundation/engine/pull/734)) Returns service sid in commands instead of hash.
- ([#724](https://github.com/mesg-foundation/engine/pull/724)) Changed system services deployment system. ([#727](https://github.com/mesg-foundation/engine/pull/727)). ([#725](https://github.com/mesg-foundation/engine/pull/725)). ([#743](https://github.com/mesg-foundation/engine/pull/743)).

#### Fixed

- ([#738](https://github.com/mesg-foundation/engine/pull/738)) Fix stream disconnection because of more than 15min of inactivity. ([#739](https://github.com/mesg-foundation/engine/pull/739)). ([#742](https://github.com/mesg-foundation/engine/pull/742)). ([#744](https://github.com/mesg-foundation/engine/pull/744)).

#### Documentation

- ([#721](https://github.com/mesg-foundation/engine/pull/721)) Move documentation to [dedicated repository](https://github.com/mesg-foundation/docs).

## [v0.7.0](https://github.com/mesg-foundation/engine/releases/tag/v0.7.0)

### [Click here to see the release notes](https://forum.mesg.com/t/mesg-core-v0-7-release-notes/197)

#### Added

- ([#677](https://github.com/mesg-foundation/engine/pull/677)) Stream acknowledgement system. The core notifies client when streams are ready.
- ([#679](https://github.com/mesg-foundation/engine/pull/679)) Add support of repeated parameters to service definition. ([#680](https://github.com/mesg-foundation/engine/pull/680)). ([#684](https://github.com/mesg-foundation/engine/pull/684)).
- ([#682](https://github.com/mesg-foundation/engine/pull/682)) Add support of type Any to service definition. ([#689](https://github.com/mesg-foundation/engine/pull/689)).
- ([#691](https://github.com/mesg-foundation/engine/pull/691)) Add database transaction mechanism to database execution.
- ([#696](https://github.com/mesg-foundation/engine/pull/696)) Add support of nested type definition for type Object.
- ([#704](https://github.com/mesg-foundation/engine/pull/704]) Move go-service to package client/service.

#### Changed

- ([#688](https://github.com/mesg-foundation/engine/pull/688)) Change sid auto-generated prefix.
- ([#699](https://github.com/mesg-foundation/engine/pull/699)) Updated to golang v1.11.4.

#### Fixed

- ([#687](https://github.com/mesg-foundation/engine/pull/687)) Fix execution generated id.
- ([#703](https://github.com/mesg-foundation/engine/pull/703)) Return error when core is not running in command dev and deploy.

#### Removed

- ([#675](https://github.com/mesg-foundation/engine/pull/675)) Remove workflow grpc client.
- ([#693](https://github.com/mesg-foundation/engine/pull/693)) Remove vendor folder.

## [v0.6.0](https://github.com/mesg-foundation/engine/releases/tag/v0.6.0)

### [Click here to see the release notes](https://forum.mesg.com/t/mesg-core-v0-6-release-notes/166)

#### Added

- ([#641](https://github.com/mesg-foundation/engine/pull/641)) Services definition accept env variables. Users can override them on deploy. [#660](https://github.com/mesg-foundation/engine/pull/660). [#666](https://github.com/mesg-foundation/engine/pull/666).
- ([#651](https://github.com/mesg-foundation/engine/pull/651)) Error added in task execution result.

#### Changed

- ([#611](https://github.com/mesg-foundation/engine/pull/611)) Switch to go1.11.
- ([#648](https://github.com/mesg-foundation/engine/pull/672)) Print all service definition in command `service detail`.
- ([#649](https://github.com/mesg-foundation/engine/pull/649)) Lowercase sid.

#### Documentation

- ([#638](https://github.com/mesg-foundation/engine/pull/638)) Fix marketplace link
- ([#643](https://github.com/mesg-foundation/engine/pull/643)) Add instruction to start the core without CLI
- ([#656](https://github.com/mesg-foundation/engine/pull/656)) Show instruction to create manually system services folder
- ([#665](https://github.com/mesg-foundation/engine/pull/665)) Add favicon

## [v0.5.0](https://github.com/mesg-foundation/engine/releases/tag/v0.5.0)

### [Click here to see the release notes](https://forum.mesg.com/t/mesg-core-v0-5-release-notes/136)

#### Breaking Changes

- ([#608](https://github.com/mesg-foundation/engine/pull/608)) Rename "command" property and add "args" property in service definition.

#### Added

- ([#583](https://github.com/mesg-foundation/engine/pull/583)) Add property Sid (Service ID) in service definition file. Allow a service to reuse the same volumes after stopping. [#627](https://github.com/mesg-foundation/engine/pull/627). [#613](https://github.com/mesg-foundation/engine/pull/613). [#619](https://github.com/mesg-foundation/engine/pull/619).

#### Changed

- ([#580](https://github.com/mesg-foundation/engine/pull/580)) Refactor package Daemon.
- ([#588](https://github.com/mesg-foundation/engine/pull/588)) Refactor tests of package container.
- ([#604](https://github.com/mesg-foundation/engine/pull/604)) Improve hash function.
- ([#609](https://github.com/mesg-foundation/engine/pull/609)) Delete all service in parallel in commands.
- ([#615](https://github.com/mesg-foundation/engine/pull/615)) Remove initialization of swarm but display useful error.
- ([#617](https://github.com/mesg-foundation/engine/pull/617)) Improve template of command service gen doc.
- ([#630](https://github.com/mesg-foundation/engine/pull/630)) Rename service id to hash.

#### Fixed

- ([#585](https://github.com/mesg-foundation/engine/pull/585)) Handle gracefully task executions without inputs.
- ([#598](https://github.com/mesg-foundation/engine/pull/598)) Start service dependencies one by one. Solve issue when dependencies request access to same resource.

#### Documentation

- ([#568](https://github.com/mesg-foundation/engine/pull/568)) Update what-is-mesg.md.
- ([#569](https://github.com/mesg-foundation/engine/pull/569)) Update README.md.
- ([#620](https://github.com/mesg-foundation/engine/pull/620)) Add docker swarm init steps to doc.

## [v0.4.0](https://github.com/mesg-foundation/engine/releases/tag/v0.4.0)

### [Click here to see the release notes](https://forum.mesg.com/t/mesg-core-v0-4-0-release-notes/116)

#### Added

- ([#534](https://github.com/mesg-foundation/engine/pull/534)) Access service dependencies based on the name of the dependency through the network.
- ([#555](https://github.com/mesg-foundation/engine/pull/555)) Add more logs on the CLI.

#### Changed

- ([#560](https://github.com/mesg-foundation/engine/pull/560)) Store executions in a database - Fix memory leak [#542](https://github.com/mesg-foundation/engine/pull/542)

#### Fixed

- ([#553](https://github.com/mesg-foundation/engine/pull/553)) UI issue with service execute command.
- ([#552](https://github.com/mesg-foundation/engine/pull/552)) Service dev command returns with an error when needed.
- ([#526](https://github.com/mesg-foundation/engine/pull/526)) Improve container deletion when a service is stopped.
- ([#524](https://github.com/mesg-foundation/engine/pull/524)) Fix sync issue on status send chans and sync issue on gRPC deploy stream sends.

## [v0.3.0](https://github.com/mesg-foundation/engine/releases/tag/v0.3.0)

### [Click here to see the release notes](https://forum.mesg.com/t/mesg-core-v0-3-0-release-notes/88)

#### Added

- ([#392](https://github.com/mesg-foundation/engine/pull/392)) **BREAKING CHANGE.** Add support for `.dockerignore`. Remove support of `.mesgignore` [#498](https://github.com/mesg-foundation/engine/pull/498).
- ([#383](https://github.com/mesg-foundation/engine/pull/383)) New API package. [#386](https://github.com/mesg-foundation/engine/pull/386). [#444](https://github.com/mesg-foundation/engine/pull/444). [#486](https://github.com/mesg-foundation/engine/pull/486). [#488](https://github.com/mesg-foundation/engine/pull/488).
- ([#409](https://github.com/mesg-foundation/engine/pull/409)) Add required validations on service's task, event and output data.
- ([#432](https://github.com/mesg-foundation/engine/pull/432)) Configuration of the CLI's output with `--no-color` and `--no-spinner` flags. Colorize JSON. [#453](https://github.com/mesg-foundation/engine/pull/453). [#480](https://github.com/mesg-foundation/engine/pull/480). [#484](https://github.com/mesg-foundation/engine/pull/484).
- ([#435](https://github.com/mesg-foundation/engine/pull/435)) Command `service logs` accepts multiple dependency filters with multiple use of `-d` flag.
- ([#478](https://github.com/mesg-foundation/engine/pull/478)) Allow multiple core to run on the same computer.
- ([#493](https://github.com/mesg-foundation/engine/pull/493)) Support numbers in service task's key, event's key and output's key
- ([#499](https://github.com/mesg-foundation/engine/pull/499)) Return service's status from API

#### Changed

- ([#371](https://github.com/mesg-foundation/engine/pull/371)) Delegate deployment of Service to Core. [#469](https://github.com/mesg-foundation/engine/pull/469).
- ([#404](https://github.com/mesg-foundation/engine/pull/404)) Change building tool.
- ([#413](https://github.com/mesg-foundation/engine/pull/413)) Improve command `service dev`. [#459](https://github.com/mesg-foundation/engine/pull/459).
- ([#417](https://github.com/mesg-foundation/engine/pull/417)) Service refactoring. [#402](https://github.com/mesg-foundation/engine/pull/402). [#414](https://github.com/mesg-foundation/engine/pull/414). [#454](https://github.com/mesg-foundation/engine/pull/454). [#458](https://github.com/mesg-foundation/engine/pull/458). [#464](https://github.com/mesg-foundation/engine/pull/464). [#472](https://github.com/mesg-foundation/engine/pull/472). [#490](https://github.com/mesg-foundation/engine/pull/490). [#491](https://github.com/mesg-foundation/engine/pull/491).
- ([#419](https://github.com/mesg-foundation/engine/pull/419)) Use Docker volumes for services. [#477](https://github.com/mesg-foundation/engine/pull/477).
- ([#427](https://github.com/mesg-foundation/engine/pull/427)) Refactor package Config
- ([#481](https://github.com/mesg-foundation/engine/pull/481)) Refactor package Database
- ([#485](https://github.com/mesg-foundation/engine/pull/485)) Improve CLI output. [#521](https://github.com/mesg-foundation/engine/pull/521).
- Tests improvements. [#381](https://github.com/mesg-foundation/engine/pull/381). [#384](https://github.com/mesg-foundation/engine/pull/384). [#391](https://github.com/mesg-foundation/engine/pull/391). [#446](https://github.com/mesg-foundation/engine/pull/446). [#447](https://github.com/mesg-foundation/engine/pull/447). [#466](https://github.com/mesg-foundation/engine/pull/466). [#489](https://github.com/mesg-foundation/engine/pull/489). [#501](https://github.com/mesg-foundation/engine/pull/501). [#504](https://github.com/mesg-foundation/engine/pull/504). [#506](https://github.com/mesg-foundation/engine/pull/506).

#### Fixed

- ([#401](https://github.com/mesg-foundation/engine/pull/401)) Gracefully stop gRPC servers.
- ([#429](https://github.com/mesg-foundation/engine/pull/429)) Fix issue when stopping services. [#505](https://github.com/mesg-foundation/engine/pull/505). [#526](https://github.com/mesg-foundation/engine/pull/526).
- ([#476](https://github.com/mesg-foundation/engine/pull/476)) Improve database error handling.
- ([#482](https://github.com/mesg-foundation/engine/pull/482)) Fix Service hash changed when fetching from git.

#### Removed

- ([#410](https://github.com/mesg-foundation/engine/pull/410)) Remove socket server in favor of the TCP server.

#### Documentation

- ([#415](https://github.com/mesg-foundation/engine/pull/415)) Added hall-of-fame to README. Thanks [sergey48k](https://github.com/sergey48k).
- ([#423](https://github.com/mesg-foundation/engine/pull/423)) Fix documentation issue.
- ([#474](https://github.com/mesg-foundation/engine/pull/474)) Documentation/update ux.
- ([#509](https://github.com/mesg-foundation/engine/pull/509)) Add forum link. [#513](https://github.com/mesg-foundation/engine/pull/513).
- ([#510](https://github.com/mesg-foundation/engine/pull/510)) Update ecosystem menu.
- ([#511](https://github.com/mesg-foundation/engine/pull/511)) Update tutorial page.
- ([#512](https://github.com/mesg-foundation/engine/pull/512)) Add sitemap.

## [v0.2.0](https://github.com/mesg-foundation/engine/releases/tag/v0.2.0)

#### Added
- ([#242](https://github.com/mesg-foundation/engine/pull/242)) Add more details in command `mesg-core service validate`
- ([#295](https://github.com/mesg-foundation/engine/pull/295)) Added more validation on the API for the data of `executeTask`, `submitResult` and `emitEvent`. Now if data doesn't match the service file, the API returns an error
- ([#302](https://github.com/mesg-foundation/engine/pull/302)) Possibility to use a config file in ~/.mesg/config.yml
- ([#303](https://github.com/mesg-foundation/engine/pull/303)) Add command `service dev` that build and run the service with the logs
- ([#303](https://github.com/mesg-foundation/engine/pull/303)) Add command `service execute` that execute a task on a service
- ([#316](https://github.com/mesg-foundation/engine/pull/316)) Delete service when stoping the `service dev` command to avoid to keep all the versions of the services.
- ([#317](https://github.com/mesg-foundation/engine/pull/317)) Add errors when trying to execute a service that is not running (nothing was happening before)
- ([#344](https://github.com/mesg-foundation/engine/pull/344)) Add `service execute --data` flag to pass arguments as key=value.
- ([#362](https://github.com/mesg-foundation/engine/pull/362)) Add `tags` list parameter for execution in order to be able to categorize and/or track a specific task execution.
- ([#362](https://github.com/mesg-foundation/engine/pull/362)) Add possibility to listen to results with the specific tag(s)

#### Changed
- ([#282](https://github.com/mesg-foundation/engine/pull/282)) Branch support added. You can now specify your branches with a `#branch` fragment at the end of your git url. E.g.: https://github.com/mesg-foundation/service-ethereum-erc20#websocket
- ([#299](https://github.com/mesg-foundation/engine/pull/299)) Add more user friendly errors when failing to connect to the Core or Docker
- ([#356](https://github.com/mesg-foundation/engine/pull/356)) Use github.com/stretchr/testify package
- ([#352](https://github.com/mesg-foundation/engine/pull/352)) Use logrus logging package

#### Fixed
- ([#358](https://github.com/mesg-foundation/engine/pull/358)) Fix goroutine leaks on api package handlers where gRPC streams are used. Handlers now doesn't block forever by exiting on context cancellation and stream.Send() errors.

#### Removed
- ([#303](https://github.com/mesg-foundation/engine/pull/303)) Deprecate command `service test` in favor of `service dev` and `service execute`

## [v0.1.0](https://github.com/mesg-foundation/engine/releases/tag/v0.1.0)

#### Added
- ([#267](https://github.com/mesg-foundation/engine/pull/267)) Add Command `service gen-doc` that generate a `README.md` in the service based on the informations of the `mesg.yml`
- ([#246](https://github.com/mesg-foundation/engine/pull/246)) Add .mesgignore to excluding file from the Docker build

#### Changed
- ([#247](https://github.com/mesg-foundation/engine/pull/247)) Update the `service init` command to have initial tasks and events
- ([#257](https://github.com/mesg-foundation/engine/pull/257)) Update the `service init` command to fetch for template based on the https://github.com/mesg-foundation/awesome/blob/master/templates.json file but also custom templates by giving the address of the template
- ([#261](https://github.com/mesg-foundation/engine/pull/261)) **BREAKING** More consistancy between the APIs, rename `taskData` into `inputData` for the `executeTask` API

#### Fixed
- ([#246](https://github.com/mesg-foundation/engine/pull/246)) Ignore files during Docker build

## [v0.1.0-beta3](https://github.com/mesg-foundation/engine/releases/tag/v0.1.0-beta3)

#### Added
- ([#246](https://github.com/mesg-foundation/engine/pull/246)) Add .mesgignore to excluding file from the Docker build

#### Changed
- ([#247](https://github.com/mesg-foundation/engine/pull/247)) Update the `service init` command to have initial tasks and events
- ([#257](https://github.com/mesg-foundation/engine/pull/257)) Update the `service init` command to fetch for template based on the https://github.com/mesg-foundation/awesome/blob/master/templates.json file but also custom templates by giving the address of the template
- ([#261](https://github.com/mesg-foundation/engine/pull/261)) **BREAKING** More consistancy between the APIs, rename `taskData` into `inputData` for the `executeTask` API

#### Fixed
- ([#246](https://github.com/mesg-foundation/engine/pull/246)) Ignore files during Docker build

## [v0.1.0-beta2](https://github.com/mesg-foundation/engine/releases/tag/v0.1.0-beta2)

#### Added
- ([#174](https://github.com/mesg-foundation/engine/pull/174)) Add CHANGELOG.md file
- ([#179](https://github.com/mesg-foundation/engine/pull/179)) Add filters for the core API
  - [API] Add `eventFilter` on `ListenEvent` API to get notification when an event with a specific name occurs
  - [API] Add `taskFilter` on `ListenResult` API to get notification when a result from a specific task occurs
  - [API] Add `outputFilter` on `ListenResult` API to get notification when a result returns a specific output
- ([#183](https://github.com/mesg-foundation/engine/pull/183)) Add a `configuration` attribute in the `mesg.yml` file to accept docker configuration for your service
- ([#187](https://github.com/mesg-foundation/engine/pull/187)) Stop all services when the MESG Core stops
- ([#190](https://github.com/mesg-foundation/engine/pull/190)) Possibility to `test` or `deploy` a service from a git or GitHub url
- ([#233](https://github.com/mesg-foundation/engine/pull/233)) Add logs in the `service test` command with service logs by default and all dependencies logs with the `--full-logs` flag
- ([#235](https://github.com/mesg-foundation/engine/pull/235)) Add `ListServices` and `GetService` APIs

#### Changed
- ([#174](https://github.com/mesg-foundation/engine/pull/174)) Update CI to build version based on tags
- ([#173](https://github.com/mesg-foundation/engine/pull/173)) Use official Docker client
- ([#175](https://github.com/mesg-foundation/engine/pull/175)) Changed the struct to use to start docker service
- ([#181](https://github.com/mesg-foundation/engine/pull/181)) MESG Core and Service start and stop functions wait for the docker container to actually run or stop.
- ([#183](https://github.com/mesg-foundation/engine/pull/183)) **BREAKING** Docker image is automatically injected in the `mesg.yml` file for your service. Now `dependencies` attribute is for extra dependencies so for most of service this is not necessary anymore.
- ([#212](https://github.com/mesg-foundation/engine/pull/212)) **BREAKING** Communication from services to core is now done through a token provided by the core
- ([#236](https://github.com/mesg-foundation/engine/pull/236)) CLI only use the API
- ([#234](https://github.com/mesg-foundation/engine/pull/234)) `service list` command now includes the status for every services

#### Fixed
- ([#179](https://github.com/mesg-foundation/engine/pull/179)) [Doc] Outdated documentation for the CLI
- ([#185](https://github.com/mesg-foundation/engine/pull/185)) Fix logs with extra characters when `mesg-core logs`


#### Removed
- ([#234](https://github.com/mesg-foundation/engine/pull/234)) Remove command `service status` in favor of `service list` command that includes status
