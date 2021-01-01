let wakeLockObj = null;
function acquireWakeup() {
  if ("keepAwake" in screen) {
    screen.keepAwake = true;
  } else if ("wakeLock" in navigator) {
    navigator.wakeLock.request("screen").then((wakeLock) => {
      wakeLockObj = wakeLock;
      wakeLockObj.addEventListener("release", () => {
        wakeLockObj = null;
      });
    });
  }
}

$(document).ready(() => {
  $(window).focus(acquireWakeup);
  acquireWakeup();
});

const RootComponent = {
	data() {
		return {
      loading: true,
      checked: false,
      item: {
        question: '',
        result: '',
      },
      showSetting: false,
      settings: {
        currentTopic: 'addition',
        addition: {
          number_of_addends: 2,
          max_sum: 20,
          min_sum: 0,
        },
        subtraction: {
          max_minuend: 20,
          number_of_subtrahends: 1,
        },
      },
		};
	},
  computed: {
    activeAddition() {
      return this.settings.currentTopic === 'addition' ? 'active' : '';
    },
    activeSubtraction() {
      return this.settings.currentTopic === 'subtraction' ? 'active' : '';
    },
    composeRequestURL() {
      return `/api/${this.settings.currentTopic}/generate`;
    },
    composeRequestData() {
      return this.settings[this.settings.currentTopic];
    }
  },
  mounted() {
    this.loadSetting();
    this.getQuestion();
  },
  methods: {
    getQuestion() {
      $.ajax({
        type: "POST",
        url: this.composeRequestURL,
        data: JSON.stringify(this.composeRequestData),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: this.handleQuestionResult,
      });
    },
    handleQuestionResult(data) {
      if (JSON.stringify(data) === JSON.stringify(this.item)) {
        this.getQuestion();
        return;
      }
      this.item = data;
      this.loading = false;
      this.checked = false;
    },
    check() {
      this.checked = true;
    },
    changeTopic(currentTopic) {
      this.setSetting({ ...this.settings, currentTopic });
      this.getQuestion();
    },
    loadSetting() {
      const loadedSettings = localStorage.settings;
      if(loadedSettings) {
        this.settings = JSON.parse(loadedSettings);
      }
    },
    setSetting(settings) {
      localStorage.settings = JSON.stringify(settings);
      this.settings = settings;
    },
  },
};
const app = Vue.createApp(RootComponent);
const vm = app.mount("#app");
