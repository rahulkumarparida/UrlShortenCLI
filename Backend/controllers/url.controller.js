require('dotenv').config()

const { Url , UrlCount } = require('../models/urlShortner.model.js')
const { generateHashId } = require('../utils/hashGenerator.utils.js')
const process = require('process')

const serverUrl = process.env.SERVER_URL 


const urlRegister = async (req,res) => {
    const body = req.body;
    const ip =await  req.ip;    
    console.log(ip);
    
    console.log(body);  
    if(!body || !body.original_url) {
        
        console.log(body);
        
        return res.status(400).json({'msg':'Url not found!!'})
    };

    try {
        const hashId = generateHashId()
        const short_url = serverUrl+`/${hashId}`
        let url = body.original_url
        if (body.original_url.slice(0,4) !== 'http') {
             url = 'https://'+body.original_url
        }

        const result =await Url.create({
            hash_id:hashId,
            ipAddress:ip,
            original_url:url,
            short_url
        })

        const countUrl = await UrlCount.create({
            ipAddress:req.ip,
            count:1,
            url:result
        })



        return res.json({message:"url successfully created",data:result})
        
    } catch (error) {
        console.log("Error while url register.");
        console.log("Error:",error)
        return error
    }
    
}



const getUrl = async (req,res) => {
    
    
    try {

        const response = await Url.find({})


        return res.json({message:"found the url" , data:response})

    } catch (error) {
        
        return error
    }
    
}



const redierectUrl = async (req,res) => {
    const hash_id  =  req.params.code;
    const ip = req.ip;
    try {
        const response = await Url.findOne({hash_id:hash_id})
        console.log(response.id);
        
        const countUrl =await UrlCount.findOneAndUpdate({url:response.id},
            {
                $addToSet: {ipAddress:ip},
                url:response,
                 $inc: { count: 1 } 
            },{new: true , upsert: true }) 

        console.log(countUrl);
        const url = response.original_url

        
        
        res.set('Location',url );

        return res.redirect(url) 

    } catch (error) {
        console.log(error);
        
        
        return error
    } 
}


const urlAnalytics = async (req,res) => {
    const hash_id  =  req.params.code;
    try {
        const response = await Url.findOne({hash_id:hash_id})
        const countData = await UrlCount.findOne({url:response.id})
        
        console.log(countData);
        
        return res.status(200).json({"analytics":countData, "metadata":response})


    } catch (error) {
        console.log(error);
        
        return error
    }
    
}


module.exports = {urlRegister , getUrl , redierectUrl , urlAnalytics}