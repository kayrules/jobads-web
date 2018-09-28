$(document).ready(function(){
  var total = 0
  $(".plus").on("click", function(e){
      e.preventDefault();
      var data = $(this).data('type');
      
      var tmp = parseInt($("#" + data).val()) + 1;
      $("#" + data).val(tmp);
      calculate();
  });

  $(".minus").on("click", function(e){
    e.preventDefault();
    var data = $(this).data('type');
    
    var tmp = parseInt($("#" + data).val()) - 1;
    if(tmp < 0) tmp = 0;
    $("#" + data).val(tmp);
    calculate();
  });

  $("#checkout").click(function(e){
    e.preventDefault();
    // calculate();
    alert("Open payment page.")
  });

  function calculate() {
    var form = $("#ajaxForm");
    var url = form.attr('action');

    var data = {
      "classic": $("#classic").val(),
      "standout": $("#standout").val(),
      "premium": $("#premium").val(),
    }

    $.ajax({
      type: "POST",
      crossDomain: true,
      url: url,
      data,
      success: function(data) {
        var amount = parseFloat(data.total / 100)
        $("#subtotal").html(amount.toFixed(2))
      }
    });
  }
});