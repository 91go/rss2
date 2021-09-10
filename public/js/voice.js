function unDecrypt(e, n) {
    if ("h" === e.substr(0, 1)) return e;
    function t(e, n, t, o) {
        var r = e,
            a = r.substring(0, n),
            i = r.substring(t);
        return a + o + i
    }
    var o = e.substring(41, 43),
        r = e.substring(46, 48),
        a = parseInt(e.substring(44, 45)),
        i = ["8", "5", "1", "7", "3", "6", "9", "0", "2", "4"],
        s = "";
    i.forEach(function(e, n) {
        a === e && (s = n)
    }),
        e = t(e, 0, 1, "h"),
        e = t(e, 41, 43, r),
        e = t(e, 46, 48, o),
        e = t(e, 44, 45, s);
    var c = "",
        l = "";
    return - 1 === e.indexOf("8.210.46.21") ? n ? (c = "http://149.129.87.151:9090/voice", l = e.substring(32)) : (e && (c = "http://149.129.87.151:9090/test"), l = e.substring(32).replace(/0/g, "1")) : n ? (c = "http://8.210.46.21:9090/voice", l = e.substring(29)) : (e && (c = "http://8.210.46.21:9090/test"), l = e.substring(29).replace(/0/g, "1")),
    c + l
}