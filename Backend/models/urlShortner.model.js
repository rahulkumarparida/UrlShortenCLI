const mongoose = require('mongoose')


const urlSchema = new mongoose.Schema(
    {
        hash_id:{
            type:String,
            required:true,
            unique:true
        },
        original_url:{
            type:String,
            required:true
        },
        short_url:{
            type:String,
            required:true
        },
        ipAddress:{
            type:String
        }
    },
    {timestamps:true}
)

const Url = mongoose.model('url',urlSchema)

const urlCountSchema = new mongoose.Schema(
    {
        count:{
            type:Number,
        },
        ipAddress:[String],
        url:{
            type: mongoose.Schema.Types.ObjectId,
            ref:'Url',
            required: true
        },
    }
);

const UrlCount = mongoose.model('urlCount',urlCountSchema);


module.exports = { Url , UrlCount };
