import { Spin, Typography } from 'antd';
import { LoadingOutlined } from '@ant-design/icons';
import { useEffect, useState } from 'react';

const { Text } = Typography;

const LoadingSequence = () => {
    const [loadingText, setLoadingText] = useState('');
    const loadingSequence = ['Loading .', 'Loading ..', 'Loading ...', 'Loading ....', 'Loading .....'];
  
    useEffect(() => {
      let currentIndex = 0;
      const intervalId = setInterval(() => {
        setLoadingText(loadingSequence[currentIndex]);
        currentIndex = (currentIndex + 1) % loadingSequence.length;
      }, 350);
  
      return () => clearInterval(intervalId);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);
  
    return <Text style={{color: 'rgb(22, 119, 255)', fontSize: 14, marginTop: '16px'}} strong>{loadingText}</Text>;
  };

const LargeLoader = () => (
    <div style={{width: '100%', display: 'flex', flexDirection: 'column', justifyContent: 'center', alignItems: 'center',  paddingTop: '60px'}}>
        <Spin indicator={<LoadingOutlined style={{ fontSize: 48 }}/>} />
        <div style={{display: 'flex', justifyContent: 'flex-start', width: '130px', paddingLeft: '34px'}}>
            {<LoadingSequence />}
        </div>
    </div>
)

export default LargeLoader;