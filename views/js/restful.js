"use strict";

/*--------------------------------------
 Utility
 -------------------------------------- */
function sendHTTPGETAjax(url, dataType, callback) {
    $.ajax({
        url: url,
        type: "GET",
        dataType: dataType,
        cache: false,
        success: function(res){
            callback(res);
        },
        error: function(){
            callback(null);
        }
    });
}

function sendHTTPPOSTAjax(url, data, dataType, callback) {
    $.ajax({
        url: url,
        type: "POST",
        data: data,
        dataType: dataType,
        processData: false,
        success: function(res){
            callback(res);
        },
        error: function(xhr, ajaxOptions, thrownError){
            console.log(xhr.status);
            console.log(thrownError);
        }
    });
}

/*--------------------------------------
 Restful API
 -------------------------------------- */
function Restful() {
}

Restful.prototype.GetTest = function () {
	sendHTTPGETAjax('/api?get/test', "json", function(res) {
		console.log(res);
	});
};

