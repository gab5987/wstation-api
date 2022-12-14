fetch('/measurements/last')
        .then(response => response.json())
        .then(data => {
            document.getElementById('totalMeasurements').innerHTML = data.split('$')[0];
            document.getElementById('lastMeasurementTime').innerHTML = getData(data.split('$')[4]);
        });

function getData(epochTime) {
    epochTime = Number(epochTime);
    return new Date(epochTime * 1000)
}