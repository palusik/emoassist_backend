var express  = require('express'),
    path     = require('path'),
    fs       = require('fs'),
    bodyParser = require('body-parser'),
    app = express(),
    expressValidator = require('express-validator');


/*Set EJS template Engine*/
app.set('views','./views');
app.set('view engine','ejs');

app.use(express.static(path.join(__dirname, 'public')));
app.use(bodyParser.urlencoded({ extended: true })); //support x-www-form-urlencoded
app.use(bodyParser.json());
app.use(expressValidator());

/*MySql connection*/
var connection  = require('express-myconnection'),
    mysql = require('mysql');

app.use(

    connection(mysql,{
        host     : 'emoassistdb',
        user     : 'root',
        password : '',
        database : 'emodb',
        ssl: {
            cert: fs.readFileSync('./certs/SSLCert.pem')
        },
        debug    : true //set true if you wanna see debug logger
    },'request')

);

app.get('/',function(req,res){
    res.send('Welcome to Emotional Assistant monitor system');
});


//RESTful route
var router = express.Router();


/*------------------------------------------------------
*  This is router middleware,invoked everytime
*  we hit url /api and anything after /api
*  like /api/user , /api/user/7
*  we can use this for doing validation,authetication
*  for every route started with /api
--------------------------------------------------------*/
router.use(function(req, res, next) {
    console.log(req.method, req.url);
    next();
});

var curut = router.route('/alerts');


//show the CRUD interface | GET
curut.get(function(req,res,next){


    req.getConnection(function(err,conn){

        if (err) return next("Cannot Connect");

        var query = conn.query('SELECT * FROM alerts',function(err,rows){

            if(err){
                console.log(err);
                return next("Mysql error, check your query");
            }

            res.render('alerts',{title:"All alerts",data:rows});

         });

    });

});




//now for Single route (GET,DELETE,PUT)
var curut2 = router.route('/alerts/:idalerts');

/*------------------------------------------------------
route.all is extremely useful. you can use it to do
stuffs for specific routes. for example you need to do
a validation everytime route /api/user/:user_id it hit.

remove curut2.all() if you dont want it
------------------------------------------------------*/
curut2.all(function(req,res,next){
    console.log("You need to smth about curut2 Route ? Do it here");
    console.log(req.params);
    next();
});

//get data to update
curut2.get(function(req,res,next){

    var idalerts = req.params.idalerts;

    req.getConnection(function(err,conn){

        if (err) return next("Cannot Connect");

        var query = conn.query("SELECT * FROM alerts WHERE idalerts = ? ",[idalerts],function(err,rows){

            if(err){
                console.log(err);
                return next("Mysql error, check your query");
            }

            //if user not found
            if(rows.length < 1)
                return res.send("Alert is Not found");

            res.render('edit',{title:"Edit Alert",data:rows});
        });

    });

});

//update data
curut2.put(function(req,res,next){
    var idalerts = req.params.idalerts;

    //get data
    var data = {
        created:req.body.created,
        userid:req.body.userid,
        type:req.body.type,
        status:req.body.status
     };

    //inserting into mysql
    req.getConnection(function (err, conn){

        if (err) return next("Cannot Connect");

        var query = conn.query("UPDATE alerts set ? WHERE idalerts = ? ",[data,idalerts], function(err, rows){

           if(err){
                console.log(err);
                return next("Mysql error, check your query");
           }

          res.sendStatus(200);

        });

     });

});

//delete data
curut2.delete(function(req,res,next){

    var idalerts = req.params.idalerts;

     req.getConnection(function (err, conn) {

        if (err) return next("Cannot Connect");

        var query = conn.query("DELETE FROM alerts  WHERE idalerts = ? ",[idalerts], function(err, rows){

             if(err){
                console.log(err);
                return next("Mysql error, check your query");
             }

             res.sendStatus(200);

        });
        //console.log(query.sql);

     });
});


// actions operations
var curut3 = router.route('/actions');


//show the CRUD interface | GET
curut3.get(function(req,res,next){


    req.getConnection(function(err,conn){

        if (err) return next("Cannot Connect");

        var query = conn.query('SELECT * FROM actions',function(err,rows){

            if(err){
                console.log(err);
                return next("Mysql error, check your query");
            }

            res.render('actions',{title:"All Actions",data:rows});

        });

    });

});


//post data to DB | POST
curut3.post(function(req,res,next){

    //get data
    var data = {
        userid:req.body.userid,
        haction:req.body.haction,
        naction:req.body.naction,
        aaction:req.body.aaction,
        saction:req.body.saction,
        faction:req.body.faction,
        call:req.body.call,
        music:req.body.music,
        activity:req.body.activity,
        book:req.body.book,
        picture:req.body.picture
    };

    //inserting into mysql
    req.getConnection(function (err, conn){

        if (err) return next("Cannot Connect");

        var query = conn.query("INSERT INTO actions set ? ",data, function(err, rows){

            if(err){
                console.log(err);
                return next("Mysql error, check your query");
            }

            res.sendStatus(200);

        });

    });

});


// user alerts operations
var curut4 = router.route('/useralerts/:userid');
//show the CRUD interface | GET
curut4.get(function(req,res,next){

    var userid = req.params.userid;

    req.getConnection(function(err,conn){

        if (err) return next("Cannot Connect");

        var query = conn.query("SELECT * FROM alerts WHERE userid = ? ORDER BY idalerts DESC",[userid],function(err,rows){

            if(err){
                console.log(err);
                return next("Mysql error, check your query");
            }

            res.render('useralerts',{title:"User alerts",data:rows});

        });

    });

});


//now for Single route (GET,DELETE,PUT)
var curut5 = router.route('/actions/:userid');

//get data to update
curut5.get(function(req,res,next){

    var userid = req.params.userid;

    req.getConnection(function(err,conn){

        if (err) return next("Cannot Connect");

        var query = conn.query("SELECT * FROM actions WHERE userid = ? ",[userid],function(err,rows){

            if(err){
                console.log(err);
                return next("Mysql error, check your query");
            }

            //if user not found
            if(rows.length < 1)
                return res.send("Actions for this user is not availabe is Not found");

            res.render('edit_actions',{title:"Edit Actions",data:rows});
        });

    });

});

//update data
curut5.put(function(req,res,next){
    var userid = req.params.userid;

    //get data
    var data = {
        haction:req.body.haction,
        naction:req.body.naction,
        aaction:req.body.aaction,
        saction:req.body.saction,
        faction:req.body.faction,
        call:req.body.call,
        music:req.body.music,
        activity:req.body.activity,
        book:req.body.book,
        picture:req.body.picture
    };

    //inserting into mysql
    req.getConnection(function (err, conn){

        if (err) return next("Cannot Connect");

        var query = conn.query("UPDATE actions set ? WHERE userid = ? ",[data,userid], function(err, rows){

            if(err){
                console.log(err);
                return next("Mysql error, check your query");
            }

            res.sendStatus(200);

        });

    });

});

//delete data
curut5.delete(function(req,res,next){

    var userid = req.params.userid;

    req.getConnection(function (err, conn) {

        if (err) return next("Cannot Connect");

        var query = conn.query("DELETE FROM actions  WHERE userid = ? ",[userid], function(err, rows){

            if(err){
                console.log(err);
                return next("Mysql error, check your query");
            }

            res.sendStatus(200);

        });
        //console.log(query.sql);

    });
});


//now for Single route (GET,DELETE,PUT)
var curut6 = router.route('/useractions/:userid');

//get data to update
curut6.get(function(req,res,next){

    var userid = req.params.userid;

    req.getConnection(function(err,conn){

        if (err) return next("Cannot Connect");

        var query = conn.query("SELECT * FROM actions WHERE userid = ? ",[userid],function(err,rows){

            if(err){
                console.log(err);
                return next("Mysql error, check your query");
            }

            //if user not found
            if(rows.length < 1)
                return res.send("Actions for this user is not availabe is Not found");

            res.render('view_actions',{title:"View User Actions",data:rows});
        });

    });

});




//now we need to apply our router here
app.use('/api', router);

var port = process.env.PORT || 8000;
//start Server
var server = app.listen(port,function(){

   console.log("Listening to port %s",server.address().port);

});
