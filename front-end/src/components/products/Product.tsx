/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useParams } from 'react-router-dom'
import { Button, Card, Typography, notification, Col, Row } from 'antd';
import { useEffect, useState } from 'react';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { formatDateTimeToYYYYMMDDHHMM } from '../../helpers/stringDateFormatting';
// import UpdateProduct from './UpdateProduct';
import CreateNote from '../notes/CreateNote';
import { NOTE_TYPE } from '../../enums/notes';
import NoteList from '../notes/Notes';
import { Product, ProductNote } from '../../interfaces';
import { useGetLoggedInUser, useGetLoggedInUserId, useGetUsers } from '../../hooks';
import { createProductNote, getAllProductNotesByProductId, getProductById } from '../../api';
import { formatNumberWithSpaces } from '../../helpers/stringFormatting';

const { Text, Link } = Typography;

const ProductDetails = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUserId = useGetLoggedInUserId();
    const loggedInUser = useGetLoggedInUser();
    const users = useGetUsers();
    const [product, setProduct] = useState<Product | null>(null);
    const [editing, setEditing] = useState(false);
    const [noteTitle, setNoteTitle] = useState('');
    const [note, setNote] = useState('');
    const [productNotes, setProductNotes] = useState<Array<ProductNote>>([]);
    const { id } =  useParams();
    const productId = id || '';

    const onHandleNoteTitleChange = (event : any) => setNoteTitle(event.target.value);
    const onHandleNoteChange = (event : any) => setNote(event.target.value);

    useEffect(() => {
        if (!product && loggedInUserId) {
            getProductById(loggedInUserId, productId).then(response => {
                if (response?.error) {
                    api.error({
                        message: `Get product failed`,
                        description: response.message,
                        placement: 'bottom',
                        duration: 1.4
                    });
                }
                setProduct(response.data);
            })
            if (productNotes && productNotes.length === 0 && loggedInUser?.id) {
                getAllProductNotesByProductId(loggedInUser.id, productId).then(response => {
                  if (!response.error && response.data) {
                    setProductNotes(response.data)
                  }
                }).catch(error => {
                  console.log('error fetching product notes', error)
                })
              }
        }
    }, [loggedInUserId]);

    const getUserName = (id : string) => `${users?.[id]?.first_name} ${users?.[id]?.last_name}`;

    const clearNoteFields = () => {
        setNoteTitle('');
        setNote('');
      }

    const onSubmitProductNote = () => {
        const user = {
            id: loggedInUserId,
            name: `${loggedInUser.firstName} ${loggedInUser.lastName}`,
            email: loggedInUser.email

        }
        createProductNote(user, productId, noteTitle, note).then((response) => {
            api.info({
            message: response.message,
            placement: "bottom",
            duration: 1.2,
            });
        }).catch(error => {
            api.error({
            message: `Error creating note`,
            description: error.toString(),
            placement: "bottom",
            duration: 1.4,
            });
        })
    }

    return (
      <Row>
         {contextHolder}
      <Col span={16}  style={{paddingRight: '8px'}}>
        <div style={{display: 'flex', justifyContent: 'flex-end', gap: '12px', marginBottom: '4px'}}>
            {editing ? (
                <Button onClick={() => setEditing(false)}>Cancel</Button>
            ) : (
                <Button onClick={() => setEditing(true)}>Edit product</Button>
            )}    
        </div>
        <Card>
        {product && (
            <div style={{display: 'flex', justifyContent: 'flex-start', gap: '20px'}}>
                <div style={{paddingRight: '24px'}}>
                    <Text strong>Product name</Text><br />
                    <Text>{product.name}</Text><br/>
                    <Text strong>Description</Text><br/>
                    <Text>{product.description}</Text><br/>
                    <Text strong>Category</Text><br/>
                    <Text>{product.category}</Text>
                </div>
                <div style={{paddingRight: '24px'}}>
                    <Text strong>Price</Text><br />
                    <Text>{formatNumberWithSpaces(product.price)} SEK</Text><br/>
                    <Text strong>Tax</Text><br />
                    <Text>{product.tax_percentage}%</Text><br/>
                </div>
                <div style={{paddingRight: '24px'}}>
                    {hasPrivilege(loggedInUser.privileges, 'user_read') && (<>
                        <Text strong>Created by</Text><br />
                        <Link href={`/user/${product.created_by}`}>{getUserName(product.created_by)}</Link><br />
                        <Text strong>Created at</Text><br />
                        <Text>{formatDateTimeToYYYYMMDDHHMM(product.created_at.toString())}</Text><br />
                        <Text strong>Updated by</Text><br />
                        <Link href={`/user/${product.updated_by}`}>{getUserName(product.updated_by)}</Link><br />
                        <Text strong>Updated at</Text><br />
                        <Text>{formatDateTimeToYYYYMMDDHHMM(product.updated_at.toString())}</Text><br />
                    </>)}
                </div>
          </div>
        )}
        </Card>
        </Col>
        <Col span={8}>
        <Card style={{ width: '100%', height: 'fit-content'}}>
          <CreateNote
            type={NOTE_TYPE.product}
            title={noteTitle}
            onTitleChange={onHandleNoteTitleChange}
            note={note}
            onNoteChange={onHandleNoteChange}
            onClearNoteFields={clearNoteFields}
            onSubmit={onSubmitProductNote}
          />
        {hasPrivilege(loggedInUser.privileges, 'note_read') && productNotes && (
          <NoteList notes={productNotes} type={NOTE_TYPE.product} userId={loggedInUser.id} />
        )}
        </Card>
        </Col>
      </Row>
      )
}

export default ProductDetails;