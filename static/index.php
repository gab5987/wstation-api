<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>WStation API - main</title>
    <link rel="stylesheet" href="./static/index.css">
</head>
<body>
    <h1> yeah, the dev got lazy on this page ik ik :( </h1>
    <p>But it doesnt really matters, the main purpose of this page is just to to show you the routes of the API</p>
    <p>So, without further ado, here they are:</p>

    <div class="list">
        <ul>
            <li> <a href="/measurements">/measurements </a> </li>
            <li> <a href="/measurements/last"> /measurements/last </a> </li>
        </ul>
    </div>
    <br/>

    <p>But, what the heck does those routes do? you may ask...</p>
    <p>Well, heres the explanation:</p>
    <div class="list">
        <ul>
            <li> /measurements -> returns all the data from the db, as well as handles new posts</li> <br/>
            <li> /measurements/last -> returns the last row of the db </li>
        </ul>
    </div>
    <br/>

    <p>The repository for this API is avaliable <a href="">here</a></p>
    <p>You can acces the admin console(and i promise its a better looking page) <a href="">here</a></p>
    <p>And thats pretty much what it is, looking forward to add more routes for this project</p>
    <br/>

    <p>
        By the time you are reading this, the api got <a id="totalMeasurements" class="indicator">...</a> different measurements
    </p>
    <p>
        And the last measurement was made at <a id="lastMeasurementTime" class="indicator">...</a>
    </p>

    <script src="./static/index.js"></script>
</body>
</html>