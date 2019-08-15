var term,
    protocol,
    socketURL,
    socket;

var terminalContainer = document.getElementById('terminal-container');

//var optionElements = {
//  cursorBlink: document.querySelector('#option-cursor-blink')
//};
//
//optionElements.cursorBlink.addEventListener('change', createTerminal);

var Query = function () {
  // This function is anonymous, is executed immediately and 
  // the return value is assigned to QueryString!
  var query_string = {};
  var query = window.location.search.substring(1);
  var vars = query.split("&");
  for (var i=0;i<vars.length;i++) {
    var pair = vars[i].split("=");
        // If first entry with this name
    if (typeof query_string[pair[0]] === "undefined") {
      query_string[pair[0]] = decodeURIComponent(pair[1]);
        // If second entry with this name
    } else if (typeof query_string[pair[0]] === "string") {
      var arr = [ query_string[pair[0]],decodeURIComponent(pair[1]) ];
      query_string[pair[0]] = arr;
        // If third or later entry with this name
    } else {
      query_string[pair[0]].push(decodeURIComponent(pair[1]));
    }
  } 
  return query_string;
}();

function httpGet(url, successCallback, errorCallback)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() { 
        if (xmlHttp.readyState != 4) 
            return;

        if (xmlHttp.status == 200)
            successCallback  && successCallback(xmlHttp.responseText);
        else 
            errorCallback  && errorCallback();
    }
    xmlHttp.open("GET", url, true); // true for asynchronous 
    xmlHttp.send(null);
}

var name = Query.name;
var namespace = Query.namespace;
var container = Query.container;
var socket = null;

createTerminal(name, namespace, container, '/bin/bash');

function createTerminal(name, namespace, container, command) {
  while (terminalContainer.children.length) {
    terminalContainer.removeChild(terminalContainer.children[0]);
  }
  term = new Terminal({
    cursorBlink:  true
  });

  term.open(terminalContainer);
  term.fit();
  term.linkify();
  term.toggleFullscreen();
  console.log(term.cols);
  console.log(term.rows);

  var protocol = (location.protocol === 'https:') ? 'wss://' : 'ws://';
  var socketURL = protocol + location.hostname + ((location.port) ? (':' + location.port) : '') + '/container/terminal/shell/ws';
  socketURL += '?name=' + name + '&namespace=' + namespace + '&container='+ container + '&cols=' + term.cols + '&rows=' + term.rows;
  socket = new ReconnectingWebSocket(socketURL);
  socket.debug = false;
  socket.timeoutInterval = 5400;
  socket.onopen = runRealTerminal;
  socket.onclose = runFakeTerminal;
  socket.onerror = runFakeTerminal;
}


function runRealTerminal() {
  term.attach(socket);
  term._initialized = true;
}

function runFakeTerminal() {
  if (term._initialized) {
    return;
  }

  term._initialized = true;

  var shellprompt = '$ ';

  term.prompt = function () {
    term.write('\r\n' + shellprompt);
  };

  term.writeln('Welcome to xterm.js');
  term.writeln('This is a local terminal emulation, without a real terminal in the back-end.');
  term.writeln('Type some keys and commands to play around.');
  term.writeln('');
  term.prompt();


  term.on('key', function (key, ev) {
    var printable = (
      !ev.altKey && !ev.altGraphKey && !ev.ctrlKey && !ev.metaKey
    );

    if (ev.keyCode == 13) {
      term.prompt();
    } else if (ev.keyCode == 8) {
      /*
     * Do not delete the prompt
     */
      if (term.x > 2) {
        term.write('\b \b');
      }
    } else if (printable) {
      term.write(key);
    }
  });

  term.on('paste', function (data, ev) {
    console.log(data)
    term.write(data);
  });
}

