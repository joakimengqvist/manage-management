import { useState } from 'react';
import { EyeInvisibleOutlined, EyeTwoTone } from '@ant-design/icons';
import { Button, Input, Space, Card, Typography, Alert } from 'antd';
import { loginAuthenticate } from '../../api/users/loginAuthenticate'
import { useDispatch } from 'react-redux';
import { authenticate } from "../../redux/userDataSlice";

const { Title } = Typography;

const LoginForm: React.FC = () => {
  const dispatch = useDispatch();
  const [userName, setUserName] = useState('');
  const [password, setPassword] = useState('');
  const [loginErrorFeedback, setLoginErrorFeedback] = useState('');

  const Login = () => {
    loginAuthenticate(userName, /* password */ ).then(response => {
      if (response.error) {
        setLoginErrorFeedback(response.message)
        return
      }
      setLoginErrorFeedback('');
      dispatch(authenticate(response))

    }).catch(error => {
      setLoginErrorFeedback(error)
    })
  }

  return (
    <Card style={{ marginTop: '180px', padding: '0px 20px 12px 8px', maxWidth: '400px', height: 'fit-content'}}>
      <Space direction="vertical">
        <Title level={3}>Login</Title>
        <Input
          placeholder="input username"
          value={userName}
          onChange={event => setUserName(event.target.value)}
          onBlur={event => setUserName(event.target.value)}
        />
        <Input.Password
          placeholder="input password"
          iconRender={(visible) => (visible ? <EyeTwoTone /> : <EyeInvisibleOutlined />)}
          value={password}
          onChange={event => setPassword(event.target.value)}
          onBlur={event => setPassword(event.target.value)}
        />
        {loginErrorFeedback && <Alert message={loginErrorFeedback.toString()} type="error" />}
        <div style={{ display: 'flex', justifyContent: 'space-between', gap: '16px', marginTop: '8px' }}>
          <Button>Forgot password</Button>
          <Button type="primary" onClick={Login}>Login</Button>
        </div>
      </Space>
    </Card>
  );
};

export default LoginForm;