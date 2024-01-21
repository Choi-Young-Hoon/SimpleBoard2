
function init() {
    webSocketInit();
}
window.onload = init;

var isChartCreated = false;

var cpuChart;
var memoryChart;
var networkSentKBytesChart;
var networkRecvKBytesChart;
var networkSentPacketChart;
var networkRecvPacketChart;
var diskChartArray = [];

var ws;
function webSocketInit() {
    ws = new WebSocket("ws://" + window.location.hostname + ":50000/ws");
    ws.onopen = function () {
        console.log("SimpleBoard2 Connected.");
    };

    ws.onmessage = function (evt) {
        var json = JSON.parse(evt.data);
        if (isChartCreated == false) {
            cpuChart    = createDoughnutChart("cpuChart", "CPU");
            memoryChart = createDoughnutChart("memoryChart", "Memory");
            networkSentKBytesChart = createLineChart("network_sent_kbyte_chart", "Network Sent Kbytes");
            networkRecvKBytesChart = createLineChart("network_recv_kbyte_chart", "Network Recv Kbytes");
            networkSentPacketChart = createLineChart("network_sent_packet_chart", "Network Sent Packet");
            networkRecvPacketChart = createLineChart("network_recv_packet_chart", "Network Recv Packet");
            createDiskCharts(json.disk_infos);
            isChartCreated = true;
        }

        drawCpuChart(json.cpu_infos);
        drawMemoryChart(json.memory_info);
        drawNetworkChart(networkRecvKBytesChart, json.network_info.bytes_recv);
        drawNetworkChart(networkSentKBytesChart, json.network_info.bytes_sent);
        drawNetworkChart(networkRecvPacketChart, json.network_info.packets_recv);
        drawNetworkChart(networkSentPacketChart, json.network_info.packets_sent);
        drawDiskChart(json.disk_infos);
    }

    ws.onclose = function () {
        console.log("SimpleBoard2 Disconnected.");
    }

    ws.onerror = function (evt) {
        console.log(evt)
    }
}

const ChartIndex = {
    CPU: 0,
    MEMORY: 1,
    NETWORK: 2,
    DISK: 3
};

function createDoughnutChart(chartCanvasId, label) {
    var ctx = document.getElementById(chartCanvasId).getContext('2d');
    var newChart = new Chart(ctx, {
        type: 'doughnut',
        data: {
            labels: ['Used', 'Free'],
            datasets: [{
                label: label,
                data: [0, 0],
                backgroundColor: [
                    'rgba(255, 99, 132, 0.5)',
                    'rgba(54, 162, 235, 0.5)'
                ],
                borderColor: [
                    'rgba(255, 99, 132, 1)',
                    'rgba(54, 162, 235, 1)'
                ],
                borderWidth: 1
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            title: {
                display: true,
                text: label,
            },
        },

        plugins: [{
            beforeDraw: function (chart) {
                var width = chart.width,
                    height = chart.height,
                    ctx = chart.ctx;

                ctx.restore();
                var fontSize = (height / 114).toFixed(2);
                ctx.font = fontSize + "em sans-serif";
                ctx.textBaseline = "middle";

                var text = chart.config.data.datasets[0].data[0] + "%",
                    textX = Math.round((width - ctx.measureText(text).width) / 2),
                    textY = height / 2 + 15;

                ctx.fillText(text, textX, textY);
                ctx.save();
            }
        }],
    });
    return newChart;
}

function createLineChart(chartCanvasId, label) {
    const ctx = document.getElementById(chartCanvasId).getContext('2d');
    var newChart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: [], // 시간 레이블
            datasets: [{
                label: label,
                data: [], // 데이터 포인트
                backgroundColor: 'rgba(0, 123, 255, 0.5)',
                borderColor: 'rgba(0, 123, 255, 1)',
                borderWidth: 1,
            }],
        },
        options: {
            scales: {
                x: {
                    type: 'time',
                    time: {
                        unit: 'second',
                        stepSize: 10
                    }
                },
                y: {
                    beginAtZero: true
                }
            },
            animation: false,
        },
    });

    return newChart;
}

function createDiskCharts(diskChartInfo) {
    var diskInfos = diskChartInfo.disk_infos;
    for (var i = 0; i < diskInfos.length; i++) {
        var newCanvasID = "disk_chart_" + diskInfos[i].mount_point;

        var newPartitionTitleDiv = document.createElement("div");
        newPartitionTitleDiv.className = "col-sm-3 col-md-3 col-lg-3 disk_info_" + i;
        var newPartitionTitle = document.createElement("h2");
        newPartitionTitle.innerHTML = diskInfos[i].mount_point;
        newPartitionTitleDiv.appendChild(newPartitionTitle);

        var newCanvasDiv = document.createElement("div");
        newCanvasDiv.className = "col-sm-3 col-md-3 col-lg-3";
        var newCanvas = document.createElement("canvas");
        newCanvas.id = newCanvasID;
        newCanvasDiv.appendChild(newCanvas);

        var rowDiv = document.querySelector(".disk_charts");
        rowDiv.appendChild(newPartitionTitleDiv);
        rowDiv.appendChild(newCanvasDiv);

        var chart = createDoughnutChart(newCanvasID, diskInfos[i].name);
        diskChartArray.push({
            chart: chart,
            chart_id: newCanvasID,
        });
    }
}

function chartUpdate(chart, usageValue, maxValue) {
    chart.data.datasets[0].data[0] = usageValue;
    chart.data.datasets[0].data[1] = maxValue - usageValue;
    chart.update();
}

function drawCpuChart(cpuInfos) {
    chartUpdate(cpuChart, cpuInfos.usage_percent, 100);
    updateCpuInfo(cpuInfos);
}

function updateCpuInfo(cpuInfos) {
    var cpuInfoDiv = document.querySelector(".cpu_info");
    var cpuInfo = cpuInfos.cpu_infos[0];
    var cpuInfoHtml = "<h2>CPU 정보</h2>";
    cpuInfoHtml += "<br>";
    cpuInfoHtml += getFontTag("red", "Cores: ") + cpuInfo.cores + " Cores";
    cpuInfoHtml += "<br>";
    cpuInfoHtml += getFontTag("red", "GHz: ") + cpuInfo.ghz + " GHz";
    cpuInfoHtml += "<br>";
    cpuInfoHtml += getFontTag("red", "Model: ") + cpuInfo.model_name;
    cpuInfoHtml += "<br>";
    cpuInfoHtml += getFontTag("red", "VenderID: ") + cpuInfo.vendor_id;
    cpuInfoHtml += "<br>";
    cpuInfoHtml += getFontTag("red", "사용률: ") + cpuInfos.usage_percent + "%";
   cpuInfoDiv.innerHTML = cpuInfoHtml;
}

function drawMemoryChart(memoryInfo) {
    chartUpdate(memoryChart, memoryInfo.used_percent, 100);
    updateMemoryInfo(memoryInfo);
}

function updateMemoryInfo(memoryInfos) {
    var memoryInfoDiv = document.querySelector(".memory_info");
    var memoryInfo = memoryInfos;
    var memoryInfoHtml = "<h2>Memory 정보</h2>";
    memoryInfoHtml += "<br>";
    memoryInfoHtml += getFontTag("red", "Total: ") + memoryInfo.total + " GBytes";
    memoryInfoHtml += "<br>";
    memoryInfoHtml += getFontTag("red", "Free: ") + memoryInfo.free + " GBytes";
    memoryInfoHtml += "<br>";
    memoryInfoHtml += getFontTag("red", "Used: ") + memoryInfo.used + " GBytes";
    memoryInfoHtml += "<br>";
    memoryInfoHtml += getFontTag("red", "사용률: ") + memoryInfo.used_percent + "%";
    memoryInfoDiv.innerHTML = memoryInfoHtml;
}

function drawNetworkChart(networkChart, value) {
    const now = new Date();
    const newData = value;

    // 데이터와 레이블 추가
    networkChart.data.labels.push(now);
    networkChart.data.datasets.forEach((dataset) => {
        dataset.data.push(newData);
    });

    // 10개 데이터를 유지
    if (networkChart.data.labels.length > 10) {
        networkChart.data.labels.shift();
        networkChart.data.datasets.forEach((dataset) => {
            dataset.data.shift();
        });
    }

    networkChart.update();
}

function drawDiskChart(diskInfo) {
    var diskInfos = diskInfo.disk_infos;
    for (var i = 0; i < diskInfos.length; i++) {
        var chart = findDiskChart("disk_chart_" + diskInfos[i].mount_point);
        if (chart == null) {
            continue;
        }
        chartUpdate(chart, diskInfos[i].usage_percent, 100);
        updateDiskInfo(i, diskInfos[i]);
    }
}

function updateDiskInfo(i, diskInfo) {
    var diskInfoDiv = document.querySelector(".disk_info_" + i);
    var diskInfoHtml = "<h2>" + diskInfo.mount_point + " 정보</h2>";
    diskInfoHtml += "<br>";
    diskInfoHtml += getFontTag("red", "Device: ") + diskInfo.device;
    diskInfoHtml += "<br>";
    diskInfoHtml += getFontTag("red", "Mount Point: ") + diskInfo.mount_point;
    diskInfoHtml += "<br>";
    diskInfoHtml += getFontTag("red", "Total: ") + diskInfo.total + " GBytes";
    diskInfoHtml += "<br>";
    diskInfoHtml += getFontTag("red", "Free: ") + diskInfo.free + " GBytes";
    diskInfoHtml += "<br>";
    diskInfoHtml += getFontTag("red", "Used: ") + diskInfo.used + " GBytes";
    diskInfoHtml += "<br>";
    diskInfoHtml += getFontTag("red", "사용률: ") + diskInfo.usage_percent + "%";
    diskInfoHtml += "<br>";
    diskInfoDiv.innerHTML = diskInfoHtml;
}

function findDiskChart(canvasID) {
    for (var i = 0; i < diskChartArray.length; i++) {
        if (diskChartArray[i].chart_id == canvasID) {
            return diskChartArray[i].chart;
        }
    }
    return null;
}

function getFontTag(color, value) {
    var fontTag = "<b><font color=\" " + color + "\"></b>" + value + "</font>";
    return fontTag;
}