(function() {
  var l = window.console.log;

  /**
   * Console type is from https://github.com/google/closure-compiler/blob/b85c29855a714cae446f61c8a7b2ccaac747722b/externs/browser/webkit_dom.js
   *
   * @this {Console}
   */
  window.console.log = function() {
    // Call the original log function.
    l.apply(this, arguments);

    var ss = [];
    for (var i = 0; i < arguments.length; ++i) {
      ss.push(JSON.stringify(arguments[i]));
    }
    var p = document.createElement('p');
    p.innerHTML = ss.join(' ');

    var d = document.getElementById('debug-console');
    if (d.children.length > 0) {
      d.appendChild(document.createElement('hr'));
    }
    d.appendChild(p);
    d.scrollTop = d.scrollHeight;
  };
})();
