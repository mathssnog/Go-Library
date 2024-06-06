// document.getElementById("userForm").addEventListener("submit", function(event){
//     event.preventDefault();
    
//     const formData = new FormData(event.target);
//     const data = {};
//     formData.forEach((value, key) => (data[key] = value));

//     fetch(event.target.action, {
//         method: "POST",
//         headers: {
//             "Content-Type": "application/json"
//         },
//         body: JSON.stringify(data)
//     })
//     .then(response => response.json())
//     .then(data => {
//         document.getElementById("response").innerHTML = `
//             <p>Name: ${data.name}</p>
//             <p>Email: ${data.email}</p>
//             <p>ID: ${data.id}</p>
//         `;
//         console.log("Success:", data);
//     })
//     .catch((error) => {
//         console.error("Error:", error);
//     });
// });
