import LoginForm from '../components/login/LoginForm';

const Login = () => {
    return (
        <div style={{width: '100%', height: '100%', display: 'flex', justifyContent: 'center', gap: '40px'}}>
            <LoginForm />
            <img height="410" src="./gopher.png" />
        </div>
    )

}

export default Login;