<script>
    let message = ""
    fetch('/ping')
        .then( async (response) => {
            let data = await response.json();
            if( response.status === 200 ){
                message = data.message
            }
        })
        .catch((err) => {
            console.log(err);
        })
</script>

<h1>Index of page1</h1>

<h1>Message from Go server: {message}</h1>