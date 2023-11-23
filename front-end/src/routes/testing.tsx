/* eslint-disable @typescript-eslint/no-explicit-any */

import { Col, Input, Row, Typography } from 'antd';
import { xml_to_json, json_to_xml } from 'shapeshift-data-formatter';

const { Text } = Typography;

const Testing = () => {
    return (<>
        <Row style={{paddingBottom: '60px'}}>
            <Col span={12} style={{padding: '20px'}}>
                    <Text>XML to JSON</Text>
                    <Input.TextArea
                        style={{width: '100%', height: '200px'}}
                        placeholder="XML"
                        onChange={(e : any) => console.log(xml_to_json(e.target.value))}
                    />
            </Col>
            <Col span={12} style={{padding: '20px'}}>
                    <Text>JSON to XML</Text>
                    <Input.TextArea
                        style={{width: '100%', height: '200px'}}
                        placeholder="JSON"
                        onChange={(e : any) => console.log(json_to_xml(e.target.value))}
                    />
            </Col>
        </Row>
    </>)

}

export default Testing;