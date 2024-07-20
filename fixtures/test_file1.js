// BestBy 1/2020
//
// best_by 07/2025

// THIS IS RANDOM CODE  FOR TESTING PURPOSES.  the goal of this file is to test the BestBy and TODO in comments
// But in case you are wondering, this is a tiny testing framework for JavaScript.
// this is my first public repository on github. 15 years ago :,)
// https://github.com/bigodines/nanoTest/tree/master
// BestBy 12/2099
var NanoTest = {
  assert: function (message, test) {
    if (!test) return message;
  },
  run: function (test) {
    var message = test();
    if (message) return message;
  },
  // TODO: water plants
  runAll: function () {
    var re = /^nanoTest_*/;
    // TODO detect this one also
    var fail = new Array();
    for (var d in window) {
      if (re.test(d)) {
        var msg = NanoTest.run(window[d]);
        if (msg) fail.push(msg);
      }
    }
    return fail;
  },
};
