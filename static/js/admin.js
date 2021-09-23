$(function(){
  $('#loader').hide()
})

function generateAPIKeys(){
  $('#loader').show()
  $.post('/api-keys', data => {
    if(data && data.data){
      $('#log').val(data.data)
    }
    $('#loader').hide()
  })
}

function displaySuccess(msg){
  swal("Success", msg, "success");
}

function displayWarning(msg){
  swal("Warning", msg, "warning");
}

function displayError(msg){
  swal("Error", msg, "error");
}

function setLoaderMessage(message){
  $('#loader-text').html(message)
}
