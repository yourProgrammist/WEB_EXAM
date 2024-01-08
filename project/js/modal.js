var modal = document.getElementsByClassName("modal")[0];
var btn = document.getElementsByClassName("bron");
var edit = document.getElementsByClassName("edit");
var close = document.getElementsByClassName("btn-close");
var cancel = document.getElementsByClassName("btn-cancel");
var send = document.getElementsByClassName("btn-send");

for (var i = 0; i < btn.length; i++) {
    btn[i].onclick = function() {
        var guid = modal.getElementsByClassName("guid-name")[0];
        var routr = modal.getElementsByClassName("route-name")[0];
        var price = modal.getElementsByClassName("price")[0];
        var data_list = this.value.split("|")

        guid.innerHTML = data_list[0];
        routr.innerHTML = data_list[1];
        price.innerHTML = this.parentNode.parentNode.getElementsByClassName("price-hour")[0].innerHTML;
        modal.style.display = "flex";
        modal.getElementsByClassName("btn-send")[0].id = "";

        getPrice(modal.getElementsByClassName("price")[0].innerHTML, modal)
    }
}

for (var i = 0; i < edit.length; i++) {
    edit[i].onclick = function() {
        clearModal()
        var guid = modal.getElementsByClassName("guid-name")[0];
        var routr = modal.getElementsByClassName("route-name")[0];
        var price = modal.getElementsByClassName("price")[0];
        var date = modal.getElementsByClassName("date-input")[0];
        var time = modal.getElementsByClassName("time-input")[0];
        var persons = modal.getElementsByClassName("peoples-input")[0];
        var timeCount = modal.getElementsByClassName("count-time-input")[0];
        var op1 = modal.getElementsByClassName("opt-1")[0];
        var op2 = modal.getElementsByClassName("opt-2")[0];
        var price = modal.getElementsByClassName("price")[0];

        guid.innerHTML = this.parentNode.getElementsByClassName("guid-name")[0].innerHTML.replaceAll("+", " ");
        routr.innerHTML = this.parentNode.getElementsByClassName("route-name")[0].innerHTML;
        price.innerHTML = this.parentNode.getElementsByClassName("price")[0].innerHTML;
        date.value = this.parentNode.getElementsByClassName("date")[0].innerHTML;
        time.value = this.parentNode.getElementsByClassName("time")[0].innerHTML;
        persons.value = this.parentNode.getElementsByClassName("persons")[0].innerHTML;
        timeCount.value = this.parentNode.getElementsByClassName("duration")[0].innerHTML;
        op1.checked = (this.parentNode.getElementsByClassName("optionFirst")[0].innerHTML === 'true');
        op2.checked = (this.parentNode.getElementsByClassName("optionSecond")[0].innerHTML === 'true');
        modal.style.display = "flex";
        modal.getElementsByClassName("btn-send")[0].id = "edit";
        modal.getElementsByClassName("btn-send")[0].value = this.id;

        getPrice(this.parentNode.getElementsByClassName("start_price")[0].innerHTML, modal)
    }
}

for (var i = 0; i < close.length; i++) {
    close[i].onclick = function() {
        modal.style.display = "none";
    }
    cancel[i].onclick = function() {
        modal.style.display = "none";
    }
}

for (var i = 0; i < send.length; i++) {
    send[i].onclick = function() {
        var mainDiv = this.parentNode.parentNode;  // modal-window
        var date = mainDiv.getElementsByClassName("date-input")[0].value;
        var guid = mainDiv.getElementsByClassName("guid-name")[0].innerHTML;
        var time = mainDiv.getElementsByClassName("time-input")[0].value;
        var counTime = mainDiv.getElementsByClassName("count-time-input")[0].value;
        var peoples = mainDiv.getElementsByClassName("peoples-input")[0].value;
        var op1 = mainDiv.getElementsByClassName("opt-1")[0].checked;
        var op2 = mainDiv.getElementsByClassName("opt-2")[0].checked;
        var routeName = mainDiv.getElementsByClassName("route-name")[0].innerHTML;
        var price = mainDiv.getElementsByClassName("price")[0].innerHTML;
        
        console.log(guid);
        if (this.id == "edit") {
            uuid = this.value
            fetch("/api/orders/" + uuid, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    "date": date,
                    "time": time,
                    "guide_id": guid,
                    "duration": counTime,
                    "persons": peoples,
                    "option_first": op1.toString(),
                    "option_second": op2.toString(),
                    "route_id": routeName,
                    "price": price,
                    "student_id": mainDiv.getElementsByClassName("price")[0].value
                })
            })
            .then(function() {
                window.location.href = "/order/";
            })
        } else {
            // send post fetch
            fetch("/api/orders/", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    "date": date,
                    "time": time,
                    "guide_id": guid,
                    "duration": counTime,
                    "persons": peoples,
                    "option_first": op1.toString(),
                    "option_second": op2.toString(),
                    "route_id": routeName,
                    "price": price,
                    "student_id": mainDiv.getElementsByClassName("price")[0].value
                })
            })
            .then(function() {
                window.location.href = "/order/";
            })
        }  
    }
}

function getPrice(startPrice, mainDiv) {
    var template = function() {
        if (this.className != "date-input" && this.className != "time-input") {
            this.value = this.value.replace(/[^0-9.]/g, '').replace(/(\..*)\./g, '$1');
            if (parseInt(this.value) > 20 && this.className ==  "peoples-input") {
                this.value = 20;
            }
        }

        var hours = parseInt(mainDiv.getElementsByClassName("count-time-input")[0].value)
        var persons = parseInt(mainDiv.getElementsByClassName("peoples-input")[0].value)
        var date = mainDiv.getElementsByClassName("date-input")[0].value
        var time = String(mainDiv.getElementsByClassName("time-input")[0].value)
        var op1 = mainDiv.getElementsByClassName("opt-1")[0]
        var op2 = mainDiv.getElementsByClassName("opt-2")[0]
        time = parseInt(time.split(':')[0]);
        var day = new Date(String(date)).getDay()
        var priseOfVisitors = 0
        var res = 1
        if (!(hours > 0)) { hours = 1; }
        
        if (day == 0 || day == 6) {
            day = 1.5
        } else {
            day = 1
        }

        if (!(persons > 0)) { 
            priseOfVisitors = 0; 
        } else if (persons > 0 && persons < 5 ) { 
            priseOfVisitors = 0
        } else if (persons >= 5 && persons < 10) {
            priseOfVisitors = 1000
        } else if (persons >= 10 && persons <= 20) {
            priseOfVisitors = 1500
        }

        if (op1.checked || op2.checked)
        {
            if (op1.checked) {
                res = res * 1.3;
            }
            if (op2.checked)
            {
                res = res * 1.5;
            }
        }

        console.log(time >= 20 && time <= 23);
        if (time >= 9 && time <= 12) {
            time = 400; 
        } else if (time >= 20 && time <= 23) {
            time = 1000;
        } else {
            time = 0;
        }
        console.log(startPrice, day, hours, priseOfVisitors, time, res);

        mainDiv.getElementsByClassName("price")[0].value = startPrice;
        mainDiv.getElementsByClassName("price")[0].innerHTML = (startPrice * day * hours + priseOfVisitors + time) * res;
    }
    mainDiv.getElementsByClassName("count-time-input")[0].oninput = template;
    mainDiv.getElementsByClassName("peoples-input")[0].oninput = template;
    mainDiv.getElementsByClassName("date-input")[0].oninput = template;
    mainDiv.getElementsByClassName("time-input")[0].oninput = template;
    mainDiv.getElementsByClassName("opt-1")[0].onchange = template;
    mainDiv.getElementsByClassName("opt-2")[0].onchange = template;
}

function clearModal() {
    var date = modal.getElementsByClassName("date-input")[0];
    var time = modal.getElementsByClassName("time-input")[0];
    var persons = modal.getElementsByClassName("peoples-input")[0];
    var timeCount = modal.getElementsByClassName("count-time-input")[0];
    var op1 = modal.getElementsByClassName("opt-1")[0];
    var op2 = modal.getElementsByClassName("opt-2")[0];

    // atter
    date.readOnly = false;
    time.readOnly = false;
    persons.readOnly = false;
    timeCount.readOnly = false;
    op1.onclick = function () { return true; };
    op2.onclick = function () { return true; };
    modal.getElementsByClassName("buttons")[0].style.display = "flex";
}
