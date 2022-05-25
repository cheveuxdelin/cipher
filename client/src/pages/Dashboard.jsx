import React from 'react'
import Morse from './ciphers/Morse';
import Caesar from './ciphers/Caesar';
import Atbash from './ciphers/Atbash';

import Box from '@mui/material/Box';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import { Button, CssBaseline } from '@mui/material';
import { AppBar } from '@mui/material';
import { Toolbar } from '@mui/material';
import { Typography } from '@mui/material';
import { Container } from '@mui/material';

import Radio from '@mui/material/Radio';
import FormControlLabel from '@mui/material/FormControlLabel';

const url = "http://localhost:8080";
const structures = {
    "Morse": {
        text: "",
    },
    "Caesar": {
        text: "",
        n: 0,
    },
    "Atbash": {
        text: "",
    }
}

export default function Dashboard(props) {
    const [method, setMethod] = React.useState("Morse");
    const [action, setAction] = React.useState("Encode");
    const [body, setBody] = React.useState(structures[method]);

    const URL = `${url}/${method.toLowerCase()}/${action.toLowerCase()}`;

    async function requestHandler() {
        const response = await fetch(URL, {
            method: "POST",
            headers: { "content-type": "application/json" },
            body: JSON.stringify(body),
        })
        const data = await response.text();

        if (!response.ok) {
            alert(data);
        } else {
            props.cipherSubmitHandler(method, action, data);
        }
    };

    function changeHandler(id, value) {
        setBody(oldState => {
            return {
                ...oldState,
                [id]: value,
            }
        })
    };

    function changeMethodHandler(event) {
        setMethod(event.target.value);
        setBody(structures[event.target.value]);
    }

    function changeActionHandler(event) {
        setAction(event.target.value);
    }

    const a = {
        "Morse": <Morse data={body} changeHandler={changeHandler} submitHandler={requestHandler} />,
        "Caesar": <Caesar data={body} changeHandler={changeHandler} submitHandler={requestHandler} />,
        "Atbash": <Atbash data={body} changeHandler={changeHandler} submitHandler={requestHandler} />,
    };


    return (
        <Box sx={{
            display: 'flex',
            height: "100vh",
        }}>
            <CssBaseline />
            <AppBar
                position="fixed"
                sx={{
                    backgroundColor: "#222"
                }}
            >
                <Toolbar sx={{
                    display: "grid",
                    gridTemplateColumns: "1fr auto 1fr",
                    alignItems: "center",
                }}>
                    <Typography variant="h6" noWrap component="div">
                        Cipher API
                    </Typography>

                    <Box>

                        <FormControl>
                            <InputLabel id="method">Method</InputLabel>
                            <Select
                                labelId="method"
                                id="Method"
                                value={method}
                                label="method"
                                onChange={changeMethodHandler}
                                size="small"
                                sx={{ color: "white", }}
                            >
                                {Object.keys(a).map(k => <MenuItem value={k}>{k}</MenuItem>)}
                            </Select>
                        </FormControl>

                        <FormControlLabel value="Encode" onChange={changeActionHandler} control={<Radio />} label="Encode" checked={action == "Encode"} />
                        <FormControlLabel value="Decode" onChange={changeActionHandler} control={<Radio />} label="Decode" checked={action == "Decode"} />
                    </Box>

                    <Button sx={{
                        width: "fit-content",
                        justifySelf: "end"
                    }}>login</Button>

                </Toolbar>
            </AppBar>
            <Box
                component="main"
                sx={{
                    width: "100%",
                    bgcolor: 'background.default',
                    display: "flex",
                    flexDirection: "column",
                    height: "100%",
                }}
            >
                <Toolbar />
                <div style={{
                    display: "flex",
                    flexGrow: "1",
                }}>
                    <Box sx={{ width: "100%", height: "100%", paddingTop: "30px" }}>
                        <Container>
                            {a[method]}
                        </Container>
                    </Box>

                    <Box sx={{
                        width: "1000px", backgroundColor: "#222", display: "flex", flexDirection: "column",
                        maxHeight: "100%",
                    }}>
                        <Box sx={{
                            fontFamily: "monospace",
                            height: "300px",
                            backgroundColor: "black",
                            wordBreak: "break-all",
                            color: "#00ff00",
                            overflowY: "scroll",
                            p: 2
                        }}>
                            curl -X POST {URL} -H {"Content-Type:application/json"} -d '{JSON.stringify(body)}'
                        </Box>
                        <Box sx={{
                            maxHeight: "500px",
                            overflowY: "scroll",
                            color: "white",
                        }}>

                            {props.values && Object.keys(props.values).reverse().map(k => <div>
                                <p>{props.values[k].action}</p>
                                <p>{props.values[k].text}</p>
                                <p>{props.values[k].method}</p>
                            </div>
                            )}

                        </Box>
                    </Box>
                </div>
            </Box>
        </Box >
    )
};
