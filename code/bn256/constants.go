                                                      
                                                     
                                                 

package bn256

import (
	"math/big"
)

func bigFromBase10(s string) *big.Int {
	n, _ := new(big.Int).SetString(s, 10)
	return n
}

                                                              
var u = bigFromBase10("4965661367192848881")

                                                                          
var P = bigFromBase10("21888242871839275222246405745257275088696311157297823662689037894645226208583")

                                                                                  
var Order = bigFromBase10("21888242871839275222246405745257275088548364400416034343698204186575808495617")

                                                   
var xiToPMinus1Over6 = &gfP2{bigFromBase10("16469823323077808223889137241176536799009286646108169935659301613961712198316"), bigFromBase10("8376118865763821496583973867626364092589906065868298776909617916018768340080")}

                                                   
var xiToPMinus1Over3 = &gfP2{bigFromBase10("10307601595873709700152284273816112264069230130616436755625194854815875713954"), bigFromBase10("21575463638280843010398324269430826099269044274347216827212613867836435027261")}

                                                   
var xiToPMinus1Over2 = &gfP2{bigFromBase10("3505843767911556378687030309984248845540243509899259641013678093033130930403"), bigFromBase10("2821565182194536844548159561693502659359617185244120367078079554186484126554")}

                                                            
var xiToPSquaredMinus1Over3 = bigFromBase10("21888242871839275220042445260109153167277707414472061641714758635765020556616")

                                                                                             
var xiTo2PSquaredMinus2Over3 = bigFromBase10("2203960485148121921418603742825762020974279258880205651966")

                                                                                         
var xiToPSquaredMinus1Over6 = bigFromBase10("21888242871839275220042445260109153167277707414472061641714758635765020556617")

                                                     
var xiTo2PMinus2Over3 = &gfP2{bigFromBase10("19937756971775647987995932169929341994314640652964949448313374472400716661030"), bigFromBase10("2581911344467009335267311115468803099551665605076196740867805258568234346338")}
