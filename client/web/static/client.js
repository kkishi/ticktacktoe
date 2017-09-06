var connection = new WebSocket(
    'ws://' + location.host.replace(":8081", ":8080") + "/ttt");
connection.onopen = function() {
    connection.send('{"join":{"name":"Tick"}}');
};
connection.error = function(error) {
    console.log('websocket error: ' + error);
    connection.close();
};
var taken = [
    [false, false, false],
    [false, false, false],
    [false, false, false],
];
connection.onmessage = function(e) {
    console.log(e);
    var r = JSON.parse(e.data)
    if (r.result.finish) {
        console.log(r.result.finish);
        connection.close();
        return;
    }
    var o = r.result.make_move.opponent;
    if (o) {
        taken[o.row || 0][o.col || 0] = true;
    }
    for (var i = 0; i < 3; ++i) {
        for (var j = 0; j < 3; ++j) {
            if (!taken[i][j]) {
                taken[i][j] = true;
                console.log(i, j);
                connection.send(JSON.stringify({
                    "move": {
                        "row": i,
                        "col": j,
                    },
                }));
                return;
            }
        }
    }
};
