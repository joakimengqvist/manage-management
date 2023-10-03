import { useState } from 'react';
import { Button, Row, Col, Card, Space } from 'antd';
import { useSelector } from 'react-redux';
import { State } from '../types/state';
import { sendEmail } from '../api/email/send';
import { testBroker } from '../api/test/testBroker'
import { getAllUsers } from '../api/users/getAll';
import { getAllProjects } from '../api/projects/getAll';
import { getAllPrivileges } from '../api/privileges/getAll';
import { cardShadow } from '../enums/styles';

const TestingEndpoints: React.FC = () => {
    const userId = useSelector((state : State) => state.user.id)
    const [sentPayload, setSentPayload] = useState('')
    const [receivedPayload, setReceivedPayload] = useState('')

    const testBrokerButton = () => {
        setSentPayload("empty")
        testBroker()
        .then(response => {
            setReceivedPayload(JSON.stringify(response, undefined, 4))
        })
        .catch(error => {
            setReceivedPayload(error)
        })
    }

    const testEmailButton = () => {  
        setSentPayload('to: you@example.com, from: me@example.com, subject: "test email", message: "Hello world"')
        sendEmail(userId, "you@example.com", "me@example.com", "test email","Hello world")
        .then(response => {
            setReceivedPayload(JSON.stringify(response, undefined, 4))
        })
        .catch(error => {
            setReceivedPayload(error)
        })
    }

    const testGetAllUsersButton = () => {
        getAllUsers(userId)
        .then(response => {
        setReceivedPayload(JSON.stringify(response, undefined, 4))
        })
        .catch(error => {
        setReceivedPayload(error)
        })
    }

    const testGetAllPrivilegesButton = () => {
        getAllPrivileges(userId)
        .then(response => {
        setReceivedPayload(JSON.stringify(response, undefined, 4))
        })
        .catch(error => {
        setReceivedPayload(error)
        })
    }

    const testGetAllProjectsButton = () => {
        getAllProjects(userId)
        .then(response => {
        setReceivedPayload(JSON.stringify(response, undefined, 4))
        })
        .catch(error => {
        setReceivedPayload(error)
        })
    }



    return (
        <>
        <Space style={{ padding: '16px', paddingBottom: '0px'}}>
            <Button onClick={testBrokerButton}>logger</Button>
            <Button onClick={testEmailButton}>Mail</Button>
            <Button onClick={testGetAllUsersButton}>Get all users</Button>
            <Button onClick={testGetAllProjectsButton}>Get all projects</Button>
            <Button onClick={testGetAllPrivilegesButton}>Get all privileges</Button>
        </Space>
        <Row gutter={16} style={{ padding: '16px'}}>
            <Col span={10}>
            <Card bordered={false} title="Sent payload" style={{boxShadow: cardShadow, borderRadius: 0}}>
                {sentPayload}
            </Card>
            </Col>
            <Col span={10}>
            <Card bordered={false} title="Recieved payload" style={{boxShadow: cardShadow, borderRadius: 0}}>
                {receivedPayload}
            </Card>
            </Col>
        </Row>
        </>
    );
};

export default TestingEndpoints;
