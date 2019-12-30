const url = '/api/addition/generate';
let currentResult = {};
const defaultRequest = {
  "number_of_addend": 2,
  "max_sum": 20,
};

function getQuestion() {
  $('#mathContent').html('');
  $('#checkButton').css('display', 'none');
  $('#nextButton').css('display', 'none');
  $('#loadingButton').css('display', 'block');

  $.ajax({
      type: "POST",
      url,
      data: JSON.stringify(defaultRequest),
      contentType: "application/json; charset=utf-8",
      dataType: "json",
      success: handleQuestion,
      failure: handleError, 
  });
}

function handleQuestion(data) {
  if(JSON.stringify(data) === JSON.stringify(currentResult)) {
    // To avoid duplicate random
    getQuestion();
    return;
  }
  currentResult = data;
  const question = `${currentResult.addends.join(' + ')} = `;
  $('#mathContent').html(question);
  $('#checkButton').css('display', 'block');
  $('#nextButton').css('display', 'none');
  $('#loadingButton').css('display', 'none');
}

function handleError(data) {
  $('#mathContent').html('Problem occurs! try again');
  $('#checkButton').css('display', 'none');
  $('#nextButton').css('display', 'block');
  $('#loadingButton').css('display', 'none');
}

function checkButtonClick() {
  const question = `${currentResult.addends.join(' + ')} = ${currentResult.sum}`;
  $('#mathContent').html(question);
  $('#checkButton').css('display', 'none');
  $('#nextButton').css('display', 'block');
}

function nextButtonClick() {
	getQuestion();
}

$(document).ready(() => {
  getQuestion();
  $('#checkButton').bind('click', checkButtonClick); 
  $('#nextButton').bind('click', nextButtonClick); 
});
