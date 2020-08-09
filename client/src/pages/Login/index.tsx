import React, { useState } from "react";
import { Button, Form, Input, Checkbox, message } from "antd";
import axios from "axios";
import "./index.scss";
import "antd/dist/antd.css";
//use tag Trans and t() method i18n
import { Trans, useTranslation } from "react-i18next";

export default function Login() {
  const { t } = useTranslation("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const key = "null"; // has 1 massage
  const loginLoading = (msg: string, time: number) => {
    message.loading({ content: t(msg), key, duration: time });
  };

  return (
    <div className="Login">
      <header className="Login-header">
        <h1>
          <Trans>Project Name</Trans>
        </h1>
        <Form name="login">
          <Form.Item
            name="username"
            rules={[
              {
                required: true,
                message: t("Please input your username") + "!",
              },
            ]}
          >
            <Input
              placeholder={t("Please input your username") + ":"}
              onChange={(e) => {
                setUsername(e.target.value);
              }}
            />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[
              {
                required: true,
                message: t("Please input your password") + "!",
              },
            ]}
          >
            <Input.Password
              placeholder={t("Please input your password") + ":"}
              onChange={(e) => {
                setPassword(e.target.value);
              }}
            />
          </Form.Item>
          <Form.Item>
            <Checkbox>
              <Trans>Remember me</Trans>
            </Checkbox>
          </Form.Item>
          <Form.Item>
            <Button
              type="primary"
              htmlType="submit"
              onClick={() => {
                // show longin status
                message.loading({ content: t("logging"), key, duration: 0 });
                if (username === "" || password === "") {
                  message.warning({
                    content: t("Please input your username or password"),
                    key,
                    duration: 3,
                  });
                  return;
                }
                axios_login(username, password)
                  .then((res) => {
                    if (res.data.code === 0) {
                      // dispatch(createToken(res.data.data.token))
                      localStorage.setItem("token", res.data.data.token);
                      message.success({
                        content: t("login successsful"),
                        key,
                        duration: 3,
                      });
                      // history.push("/user/home")
                    } else {
                      message.warning({
                        content: t("Incorrect username or password"),
                        key,
                        duration: 3,
                      });
                    }
                  })
                  .catch((err) => {
                    console.log(err);
                    loginLoading(err.message, 3);
                  });
              }}
            >
              <Trans>Login</Trans>
            </Button>
          </Form.Item>
        </Form>
      </header>
    </div>
  );
}
/* API document:
 *  request:
 *   username    string    phone,username,showname(utf-8),email
 *   password    string    md5+sal
 *   code        string    captcha
 *   platform_id int       Single sign on ID
 * use platform_id = 0 development platform
 * in url:"/graphql" query doc { dictionary } get more infomation
 *
 * response:
 *   code  number statu code (See dictionary table)
 *   data  playload object interface type any
 *   msg   string debug logs
 * loginResponse:
 *   data:{
 *     token string jwt-token
 *     user  object user struct
 *   }
 */

function axios_login(username: string, password: string) {
  return axios.post("/login", {
    username: username,
    password: password,
  });
}
