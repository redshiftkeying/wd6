import React from 'react'
import { Result, Button } from 'antd';
import { useHistory } from 'react-router-dom'
import { Trans, useTranslation } from 'react-i18next'
import 'antd/dist/antd.css'

export default function Miss() {
    const history = useHistory()
    const {t} =useTranslation()
    return <Result
        status="404"
        title="404"
        subTitle={t('Sorry, the page you visited does not exist')}
        extra={<Button type="primary" onClick={() => {
            history.push("/")
        }}><Trans>Back Home</Trans></Button>}
    />
}