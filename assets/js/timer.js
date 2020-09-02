var interval = 1000;
var endtime = document.getElementById('startTime').innerText;
var endDate = new Date(endtime);
var counter = 0;

var addZero = function(n){return('0'+n).slice(-2);}
var addZeroDay = function(n){return('0'+n).slice(-3);}
function countdown(period){
    var day = Math.floor(period / (1000 * 60 * 60 * 24));
    period -=  (day　*(1000 * 60 * 60 * 24));
    var hour = Math.floor(period / (1000 * 60 * 60));
    period -= (hour *(1000 * 60 * 60));
    var minutes =  Math.floor(period / (1000 * 60));
    period -= (minutes * (1000 * 60));
    var second = Math.floor(period / 1000);
    var insert = "";
    insert += '<span class="h">' + addZeroDay(day) +'day' + '</span>';
    insert += '<span class="h">' + addZero(hour) + ':'+'</span>';
    insert +=  '<span class="m">' + addZero(minutes) +':' + '</span>';
    insert += '<span class="s">' + addZero(second)+ ':'+ '</span>';
    document.getElementById('timer').innerHTML = insert;
}

function countup(){
    counter++;
    var insert = "";
    var tmp = counter
    var hour = Math.floor(tmp / (60 * 60));
    tmp -= (hour *(60 * 60));
    var minutes =  Math.floor(tmp / (60));
    tmp -= (minutes * (60));
    //var second = Math.floor(tmp / 1000);
    var second  = tmp
    //console.log(second)
    insert += '<span class="h">' + addZero(hour) + ':'+'</span>';
    insert +=  '<span class="m">' + addZero(minutes) +':' + '</span>';
    insert += '<span class="s">' + addZero(second)+ ''+ '</span>';
    document.getElementById('timer').innerHTML = insert;
}

function resetCounter(){
    counter = -3;
}
var execTimer = function(){
    var nowDate = new Date();
    var period = endDate - nowDate ;
    if(period >= 0) {
        countdown(period);
    }else{
        countup();
    }
}


setInterval(execTimer, 1000);