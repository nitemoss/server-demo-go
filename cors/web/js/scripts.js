async function request(url, data, method) {
	document.getElementById("resp").innerHTML = "Loading...";
	console.log(method)
		try {
			let res = await fetch(url, {
				method: method, // *GET, POST, PUT, DELETE, etc.
				mode: 'cors', // no-cors, *cors, same-origin
				cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
				credentials: 'same-origin', // include, *same-origin, omit
				headers: {
				  'Content-Type': 'application/json'
				  // 'Content-Type': 'application/x-www-form-urlencoded',
				},
				redirect: 'follow', // manual, *follow, error
				referrerPolicy: 'no-referrer', // no-referrer, *client
				body: data // body data type must match "Content-Type" header
			  });
			let inf = await res;
			console.log(inf);

		let jso = await res.text();
		document.getElementById("resp").innerHTML = jso;
	} catch(e){

		document.getElementById("resp").innerHTML = e;
		throw e
	}

}

function pop() {

	request(document.getElementById("url").value, document.getElementById("body").value, document.getElementById("method").value);
}
