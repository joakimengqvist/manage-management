import LoginForm from '../components/login/LoginForm';
import { Row, Col } from 'antd';

const Login: React.FC = () => {
    return (
        <div style={{padding: '12px 8px'}}>
            <Row>
            <Col span={8}>
                    <div style={{padding: '4px', maxWidth: '400px'}}>
                        <LoginForm />
                    </div>
                </Col>
                <Col span={16}>
                </Col>
            </Row>
        </div>
    )

}

export default Login;