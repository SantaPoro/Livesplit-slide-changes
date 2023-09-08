
const net = require('net');
let currentSplit = null;

(async function () {
	const client = await createClient();
	while (true) {
		await new Promise(res => setTimeout(res, 100))
		const splitIndex = await getSplitIndex(client);
		if (currentSplit + 1 == splitIndex) {
			console.log('ändra googlellgell slides')
		}
		currentSplit = splitIndex;
	}

})()

function createClient() {
	const client = new net.Socket();
	client.setEncoding('utf-8');

	return new Promise(resolve => {
		client.connect(6969, '127.0.0.1', () => resolve(client));
	})
}


function getSplitIndex(client) {
	return new Promise (resolve => {
		client.write('getsplitindex\r\n');
		function onData (data) {
			client.off('data', onData);	
			resolve(parseInt(data,10));
		}
		client.on('data', onData);
	})
}