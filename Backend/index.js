require('dotenv').config()

const express = require('express')
const mongoose = require('mongoose')
const process = require('process')
const {urlRouter ,redierectRouter} = require('./routers/url.routes.js')

const PORT = process.env.PORT
const mongo_url = process.env.MONGO_URL
const app = express()

mongoose.connect(mongo_url+'/url-shortner').then((result) => {
    console.log("Connected to MongoDB successfully.");
    
}).catch((err) => {
    console.log("Some error while connecting to the Database");
    
});
app.use(express.json());
app.use(express.urlencoded({ extended: true })); 

app.use("/api/v1" , urlRouter)
app.use('',redierectRouter)


app.get('/',(req,res)=>{
    console.log("Working");
    return res.json({msg:"Healthy,Up and Running."})
})

app.listen(PORT,(req,res)=> console.log('Srever Started!!'))
