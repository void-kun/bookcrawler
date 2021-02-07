import express from 'express';
import cookieParser from 'cookie-parser';
import logger from 'morgan';


var app = express();

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());

app.use('/', (req, res) => {
    res.json({say: "hello world"});
});

export default app;

