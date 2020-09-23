<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no" />

    <meta http-equiv="Cache-Control" content="no-cache, no-store, must-revalidate" />
    <meta http-equiv="Pragma" content="no-cache" />
    <meta http-equiv="Expires" content="0" />

    <title>Hello Kubernetes</title>

    <style>
        html,
        body {
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            font-family: monospace;
        }

        html {
            height: 100%%;
        }

        body {
            flex: 1;
            margin: 32px;

        }

        #logo {
            height: 150px;
        }

        #message {
            margin-top: 32px;
            font-size: 42px;
            color: #326ce5;
        }

        p {
            color: #333;
            text-align: center;
            margin-top: 32px;
            font-size: 24px;
        }
    </style>
</head>

<body>
    {{{ logo }}}
    <div id="message">
        Hello Kubernetes
    </div>
    <p>
        {{ hostname }}
    </p>
</body>

</html>