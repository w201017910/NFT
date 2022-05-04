var judge=true;
function change(){
    var a=$("#a").html()
    $("#a").html($("#b").html())
    $("#b").html(a)
    judge=!judge
}

var accounts = [];
var web3 = new Web3(Web3.givenProvider || "ws://localhost:7545");
var mycontract = new web3.eth.Contract(uniswap, "0x1F891dD1fdD7b6659bC70eE9A12F3a06E4f6aFC6");
var mycontract1 = new web3.eth.Contract(erc20, "0x966F1E31f406042C3c50309Dd744d3FdCc6ffBb4");
async function getAccount() {
    accounts = await ethereum.request({ method: 'eth_requestAccounts' });

    console.log(accounts[0])
}

ethereum.on('accountsChanged', function (_accounts) {
    // Time to reload your interface with accounts[0]!
    accounts = _accounts;

    console.log(_accounts[0]);

});
var token;
var _invariant;
var eth;

function balance() {
    $.ajax({
        url: '/balance',
        type: 'get',
        data:{

        },
        success: function(data){
            eth=data.eth
            token=data.token
            _invariant=web3.utils.toBN(eth).mul(web3.utils.toBN(token)).toString()
            $(".wtt").html("余额："+(token/(10**18)).toFixed(2)+"WTT")
            $(".eth").html("余额："+(eth/(10**18)).toFixed(2)+"ETH")
            $(".eth1").html((token/eth).toFixed(2)+"WTT/ETH")
            $(".wtt1").html((eth/token).toFixed(2)+"ETH/WTT")
        }
    });


}

balance()
function calculator(tt){
    if(tt==0){$(".wtt2").val(0);return;}
var test=web3.utils.toBN(tt*10**18)
    test=test.add(web3.utils.toBN(eth))
    test=test.sub(test.div(web3.utils.toBN(500)))
    test=web3.utils.toBN(token).sub(web3.utils.toBN(_invariant).div(test))
$(".wtt2").val((test.toString())/(10**18))
}
function calculator1(tt){
    if(tt==0){$(".eth2").val(0);return;}
    var test=web3.utils.toBN(tt*10**18)
    test=test.add(web3.utils.toBN(token))
    test=test.sub(test.div(web3.utils.toBN(500)))
    test=web3.utils.toBN(eth).sub(web3.utils.toBN(_invariant).div(test))
    $(".eth2").val((test.toString())/(10**18))
}
 function _sawp() {
     if (judge){
         swap1()
     }else {
         swap2()
     }
}
async function swap1(){
    await mycontract.methods.ethToTokenSwap(1,99999999999999).send({from:accounts[0],value:web3.utils.toBN($(".eth2").val()*10**18).toString()})
    balance()
}
async function swap2(){
    allowance=await mycontract1.methods.allowance(accounts[0],"0x1F891dD1fdD7b6659bC70eE9A12F3a06E4f6aFC6").call()
    if (!(allowance>=web3.utils.toBN($(".wtt2").val()*10**18).toString())){
        await mycontract1.methods.approve("0x1F891dD1fdD7b6659bC70eE9A12F3a06E4f6aFC6",web3.utils.toBN($(".wtt2").val()*10**19).toString()).send({from:accounts[0]})
    }

    await mycontract.methods.tokenToEthSwap(web3.utils.toBN($(".wtt2").val()*10**18).toString(),1,99999999999999).send({from:accounts[0]})
    balance()
}
function swaps() {
    if (judge){
        swap3()
    }else {
        swap4()
    }
}
function swap3() {
    $.ajax({
        url: '/ethToTokenSwap',
        type: 'post',
        data:{
            "value":web3.utils.toBN($(".eth2").val()*10**18).toString(),
        },
        success: function(data){
            balance()

            alert(data.value)
        }
    })


}
function swap4() {
    $.ajax({
        url: '/tokenToEthSwap',
        type: 'post',
        data:{
            "value":web3.utils.toBN($(".wtt2").val()*10**18).toString(),
        },
        success: function(data){
            balance()
            alert(data.value)
        }
    })


}