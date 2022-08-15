$(function(){
  $('#loader').hide()
})

function generateAPIKeys(){
  $('#loader').show()
  $.post('/api-keys', data => {
    if(data && data.data){
	document.getElementById('log').value = data.data.replaceAll('\\n', '\n')
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
