function update() {
    $.getJSON('/Items', function (data) {
        const table = $("#tablebody-items");
        table.empty();

        data.forEach(function (item) {
            const row = $("<tr>");

            row.append($("<td>" + item.name + "</td>"))
                .append($("<td>" + item.artist + "</td>"))
                .append($("<td>" + item.year + "</td>"))
                .append($("<td>" + item.genre + "</td>"))
                .append($("<td>" + item.available + "</td>"));

            table.append(row);
        });
    });
}

setInterval("update()", 1000)