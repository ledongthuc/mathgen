let currentResult = {};
let currentSetting;

const defaultTopic = 'addition'
const defaultRequestData = {
  "number_of_addends": 2,
  "max_sum": 20,
};

function composeRequestURL() {
  const topic = currentSetting ? currentSetting.topic : defaultTopic;

  return `/api/${topic}/generate`;
}

function composeRequestData() {
  if(!currentSetting) return defaultRequestData;
  switch(currentSetting.topic) {
    case 'addition': return defaultRequestData;
    case 'subtraction': return {
        max_minuend: 20,
        number_of_subtrahends: 1,
    }
  }
}

function getQuestion() {
  $('#mathContent').html('');
  $('#checkButton').attr("disabled");
  $('#nextButton').attr("disabled");

  $.ajax({
      type: "POST",
      url: composeRequestURL(),
      data: JSON.stringify(composeRequestData()),
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
  const question = currentResult.question;
  $('#mathContent').html(question);
  $('#checkButton').css('display', 'block').removeAttr('disabled');
  $('#nextButton').css('display', 'none').removeAttr('disabled');
}

function handleError(data) {
  $('#mathContent').html('Problem occurs! try again');
  $('#checkButton').css('display', 'none').removeAttr('disabled');
  $('#nextButton').css('display', 'block').removeAttr('disabled');
}

function checkButtonClick() {
  const result = currentResult.result;
  $('#mathContent').html(result);
  $('#checkButton').css('display', 'none');
  $('#nextButton').css('display', 'block');
}

function nextButtonClick() {
  getQuestion();
}

function settingButtonClick() {
  $('#math').css('display', 'none');
  $('#setting').css('display', 'block');
}

function backButtonClick() {
  $('#math').css('display', 'block');
  $('#setting').css('display', 'none');
}

function loadSetting() {
  let topic = window.localStorage.getItem('topic');
  topic = topic ? topic : defaultTopic;
  return {
    topic,
  }
}

function setSetting(setting) {
  window.localStorage.setItem('topic', setting.topic);
  currentSetting = setting;
  updateTopicStatus(setting);
  getQuestion();
}

function updateTopicStatus(setting) {
  $(".topic").removeClass('active').each(function(index, topicElement) {
    const topic = $(this).attr('topic');
    const currentTopic = setting.topic;
      if(topic === currentTopic) {
        $(this).addClass('active')
      }
  });
}

function topicButtonClick() {
  const topic = $(this).attr('topic');
  setSetting({ ...currentSetting, topic })
}

let wakeLockObj = null;
function acquireWakeup() {
  if ('keepAwake' in screen) {
      screen.keepAwake = true;
      console.log('screen.keepAwake is acquired');
  } else if ('wakeLock' in navigator) {
    navigator.wakeLock.request('screen').then((wakeLock) => {
      wakeLockObj = wakeLock;
      wakeLockObj.addEventListener('release', () => {
        console.log('wakeLockObj is released');
        wakeLockObj = null;
      });
      onsole.log('wakeLockObj is acquired');
    }).catch((err) => {
      console.error(err);
    })
  } else {
    console.log('Can\'t acquire wake lock')
  }
}

$(document).ready(() => {
  currentSetting = loadSetting();
  updateTopicStatus(currentSetting);
  getQuestion();
  $('#checkButton').bind('click', checkButtonClick); 
  $('#nextButton').bind('click', nextButtonClick); 
  $('#settingButton').bind('click', settingButtonClick); 
  $('#backButton').bind('click', backButtonClick); 
  $('.topic').bind('click', topicButtonClick);
  $(window).focus(acquireWakeup);
  acquireWakeup();
});
