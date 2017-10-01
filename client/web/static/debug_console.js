(function() {
  var l = window.console.log;
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
