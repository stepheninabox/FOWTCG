// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package mysqldef

const (
	// DefaultMySQLState is default state of the mySQL
	DefaultMySQLState = "HY000"
)

// MySQLState maps error code to MySQL SQLSTATE value.
// The values are taken from ANSI SQL and ODBC and are more standardized.
var MySQLState = map[uint16]string{
	ErDupKey:                              "23000",
	ErOutofmemory:                         "HY001",
	ErOutOfSortmemory:                     "HY001",
	ErConCountError:                       "08004",
	ErBadHostError:                        "08S01",
	ErHandshakeError:                      "08S01",
	ErDbaccessDeniedError:                 "42000",
	ErAccessDeniedError:                   "28000",
	ErNoDbError:                           "3D000",
	ErUnknownComError:                     "08S01",
	ErBadNullError:                        "23000",
	ErBadDbError:                          "42000",
	ErTableExistsError:                    "42S01",
	ErBadTableError:                       "42S02",
	ErNonUniqError:                        "23000",
	ErServerShutdown:                      "08S01",
	ErBadFieldError:                       "42S22",
	ErWrongFieldWithGroup:                 "42000",
	ErWrongSumSelect:                      "42000",
	ErWrongGroupField:                     "42000",
	ErWrongValueCount:                     "21S01",
	ErTooLongIdent:                        "42000",
	ErDupFieldname:                        "42S21",
	ErDupKeyname:                          "42000",
	ErDupEntry:                            "23000",
	ErWrongFieldSpec:                      "42000",
	ErParseError:                          "42000",
	ErEmptyQuery:                          "42000",
	ErNonuniqTable:                        "42000",
	ErInvalidDefault:                      "42000",
	ErMultiplePriKey:                      "42000",
	ErTooManyKeys:                         "42000",
	ErTooManyKeyParts:                     "42000",
	ErTooLongKey:                          "42000",
	ErKeyColumnDoesNotExits:               "42000",
	ErBlobUsedAsKey:                       "42000",
	ErTooBigFieldlength:                   "42000",
	ErWrongAutoKey:                        "42000",
	ErForcingClose:                        "08S01",
	ErIpsockError:                         "08S01",
	ErNoSuchIndex:                         "42S12",
	ErWrongFieldTerminators:               "42000",
	ErBlobsAndNoTerminated:                "42000",
	ErCantRemoveAllFields:                 "42000",
	ErCantDropFieldOrKey:                  "42000",
	ErBlobCantHaveDefault:                 "42000",
	ErWrongDbName:                         "42000",
	ErWrongTableName:                      "42000",
	ErTooBigSelect:                        "42000",
	ErUnknownProcedure:                    "42000",
	ErWrongParamcountToProcedure:          "42000",
	ErUnknownTable:                        "42S02",
	ErFieldSpecifiedTwice:                 "42000",
	ErUnsupportedExtension:                "42000",
	ErTableMustHaveColumns:                "42000",
	ErUnknownCharacterSet:                 "42000",
	ErTooBigRowsize:                       "42000",
	ErWrongOuterJoin:                      "42000",
	ErNullColumnInIndex:                   "42000",
	ErPasswordAnonymousUser:               "42000",
	ErPasswordNotAllowed:                  "42000",
	ErPasswordNoMatch:                     "42000",
	ErWrongValueCountOnRow:                "21S01",
	ErInvalidUseOfNull:                    "22004",
	ErRegexpError:                         "42000",
	ErMixOfGroupFuncAndFields:             "42000",
	ErNonexistingGrant:                    "42000",
	ErTableaccessDeniedError:              "42000",
	ErColumnaccessDeniedError:             "42000",
	ErIllegalGrantForTable:                "42000",
	ErGrantWrongHostOrUser:                "42000",
	ErNoSuchTable:                         "42S02",
	ErNonexistingTableGrant:               "42000",
	ErNotAllowedCommand:                   "42000",
	ErSyntaxError:                         "42000",
	ErAbortingConnection:                  "08S01",
	ErNetPacketTooLarge:                   "08S01",
	ErNetReadErrorFromPipe:                "08S01",
	ErNetFcntlError:                       "08S01",
	ErNetPacketsOutOfOrder:                "08S01",
	ErNetUncompressError:                  "08S01",
	ErNetReadError:                        "08S01",
	ErNetReadInterrupted:                  "08S01",
	ErNetErrorOnWrite:                     "08S01",
	ErNetWriteInterrupted:                 "08S01",
	ErTooLongString:                       "42000",
	ErTableCantHandleBlob:                 "42000",
	ErTableCantHandleAutoIncrement:        "42000",
	ErWrongColumnName:                     "42000",
	ErWrongKeyColumn:                      "42000",
	ErDupUnique:                           "23000",
	ErBlobKeyWithoutLength:                "42000",
	ErPrimaryCantHaveNull:                 "42000",
	ErTooManyRows:                         "42000",
	ErRequiresPrimaryKey:                  "42000",
	ErKeyDoesNotExits:                     "42000",
	ErCheckNoSuchTable:                    "42000",
	ErCheckNotImplemented:                 "42000",
	ErCantDoThisDuringAnTransaction:       "25000",
	ErNewAbortingConnection:               "08S01",
	ErMasterNetRead:                       "08S01",
	ErMasterNetWrite:                      "08S01",
	ErTooManyUserConnections:              "42000",
	ErReadOnlyTransaction:                 "25000",
	ErNoPermissionToCreateUser:            "42000",
	ErLockDeadlock:                        "40001",
	ErNoReferencedRow:                     "23000",
	ErRowIsReferenced:                     "23000",
	ErConnectToMaster:                     "08S01",
	ErWrongNumberOfColumnsInSelect:        "21000",
	ErUserLimitReached:                    "42000",
	ErSpecificAccessDeniedError:           "42000",
	ErNoDefault:                           "42000",
	ErWrongValueForVar:                    "42000",
	ErWrongTypeForVar:                     "42000",
	ErCantUseOptionHere:                   "42000",
	ErNotSupportedYet:                     "42000",
	ErWrongFkDef:                          "42000",
	ErOperandColumns:                      "21000",
	ErSubqueryNo1Row:                      "21000",
	ErIllegalReference:                    "42S22",
	ErDerivedMustHaveAlias:                "42000",
	ErSelectReduced:                       "01000",
	ErTablenameNotAllowedHere:             "42000",
	ErNotSupportedAuthMode:                "08004",
	ErSpatialCantHaveNull:                 "42000",
	ErCollationCharsetMismatch:            "42000",
	ErWarnTooFewRecords:                   "01000",
	ErWarnTooManyRecords:                  "01000",
	ErWarnNullToNotnull:                   "22004",
	ErWarnDataOutOfRange:                  "22003",
	WarnDataTruncated:                     "01000",
	ErWrongNameForIndex:                   "42000",
	ErWrongNameForCatalog:                 "42000",
	ErUnknownStorageEngine:                "42000",
	ErTruncatedWrongValue:                 "22007",
	ErSpNoRecursiveCreate:                 "2F003",
	ErSpAlreadyExists:                     "42000",
	ErSpDoesNotExist:                      "42000",
	ErSpLilabelMismatch:                   "42000",
	ErSpLabelRedefine:                     "42000",
	ErSpLabelMismatch:                     "42000",
	ErSpUninitVar:                         "01000",
	ErSpBadselect:                         "0A000",
	ErSpBadreturn:                         "42000",
	ErSpBadstatement:                      "0A000",
	ErUpdateLogDeprecatedIgnored:          "42000",
	ErUpdateLogDeprecatedTranslated:       "42000",
	ErQueryInterrupted:                    "70100",
	ErSpWrongNoOfArgs:                     "42000",
	ErSpCondMismatch:                      "42000",
	ErSpNoreturn:                          "42000",
	ErSpNoreturnend:                       "2F005",
	ErSpBadCursorQuery:                    "42000",
	ErSpBadCursorSelect:                   "42000",
	ErSpCursorMismatch:                    "42000",
	ErSpCursorAlreadyOpen:                 "24000",
	ErSpCursorNotOpen:                     "24000",
	ErSpUndeclaredVar:                     "42000",
	ErSpFetchNoData:                       "02000",
	ErSpDupParam:                          "42000",
	ErSpDupVar:                            "42000",
	ErSpDupCond:                           "42000",
	ErSpDupCurs:                           "42000",
	ErSpSubselectNyi:                      "0A000",
	ErStmtNotAllowedInSfOrTrg:             "0A000",
	ErSpVarcondAfterCurshndlr:             "42000",
	ErSpCursorAfterHandler:                "42000",
	ErSpCaseNotFound:                      "20000",
	ErDivisionByZero:                      "22012",
	ErIllegalValueForType:                 "22007",
	ErProcaccessDeniedError:               "42000",
	ErXaerNota:                            "XAE04",
	ErXaerInval:                           "XAE05",
	ErXaerRmfail:                          "XAE07",
	ErXaerOutside:                         "XAE09",
	ErXaerRmerr:                           "XAE03",
	ErXaRbrollback:                        "XA100",
	ErNonexistingProcGrant:                "42000",
	ErDataTooLong:                         "22001",
	ErSpBadSQLstate:                       "42000",
	ErCantCreateUserWithGrant:             "42000",
	ErSpDupHandler:                        "42000",
	ErSpNotVarArg:                         "42000",
	ErSpNoRetset:                          "0A000",
	ErCantCreateGeometryObject:            "22003",
	ErTooBigScale:                         "42000",
	ErTooBigPrecision:                     "42000",
	ErMBiggerThanD:                        "42000",
	ErTooLongBody:                         "42000",
	ErTooBigDisplaywidth:                  "42000",
	ErXaerDupid:                           "XAE08",
	ErDatetimeFunctionOverflow:            "22008",
	ErRowIsReferenced2:                    "23000",
	ErNoReferencedRow2:                    "23000",
	ErSpBadVarShadow:                      "42000",
	ErSpWrongName:                         "42000",
	ErSpNoAggregate:                       "42000",
	ErMaxPreparedStmtCountReached:         "42000",
	ErNonGroupingFieldUsed:                "42000",
	ErForeignDuplicateKeyOldUnused:        "23000",
	ErCantChangeTxCharacteristics:         "25001",
	ErWrongParamcountToNativeFct:          "42000",
	ErWrongParametersToNativeFct:          "42000",
	ErWrongParametersToStoredFct:          "42000",
	ErDupEntryWithKeyName:                 "23000",
	ErXaRbtimeout:                         "XA106",
	ErXaRbdeadlock:                        "XA102",
	ErFuncInexistentNameCollision:         "42000",
	ErDupSignalSet:                        "42000",
	ErSignalWarn:                          "01000",
	ErSignalNotFound:                      "02000",
	ErSignalException:                     "HY000",
	ErResignalWithoutActiveHandler:        "0K000",
	ErSpatialMustHaveGeomCol:              "42000",
	ErDataOutOfRange:                      "22003",
	ErAccessDeniedNoPasswordError:         "28000",
	ErTruncateIllegalFk:                   "42000",
	ErDaInvalidConditionNumber:            "35000",
	ErForeignDuplicateKeyWithChildInfo:    "23000",
	ErForeignDuplicateKeyWithoutChildInfo: "23000",
	ErCantExecuteInReadOnlyTransaction:    "25006",
	ErAlterOperationNotSupported:          "0A000",
	ErAlterOperationNotSupportedReason:    "0A000",
	ErDupUnknownInIndex:                   "23000",
}
