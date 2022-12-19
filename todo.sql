PGDMP     /    8                z            todo    11.18    11.18 &    -           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            .           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            /           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                       false            0           1262    16631    todo    DATABASE     �   CREATE DATABASE todo WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'English_Indonesia.1252' LC_CTYPE = 'English_Indonesia.1252';
    DROP DATABASE todo;
             postgres    false            �            1259    16861    category    TABLE     �   CREATE TABLE public.category (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at character varying(255),
    updated_at character varying(255)
);
    DROP TABLE public.category;
       public         postgres    false            �            1259    16859    category_id_seq    SEQUENCE     �   CREATE SEQUENCE public.category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.category_id_seq;
       public       postgres    false    204            1           0    0    category_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.category_id_seq OWNED BY public.category.id;
            public       postgres    false    203            �            1259    16818    gorp_migrations    TABLE     g   CREATE TABLE public.gorp_migrations (
    id text NOT NULL,
    applied_at timestamp with time zone
);
 #   DROP TABLE public.gorp_migrations;
       public         postgres    false            �            1259    16850    status_task    TABLE     �   CREATE TABLE public.status_task (
    id integer NOT NULL,
    status character varying(255) NOT NULL,
    created_at character varying(255),
    updated_at character varying(255)
);
    DROP TABLE public.status_task;
       public         postgres    false            �            1259    16848    status_task_id_seq    SEQUENCE     �   CREATE SEQUENCE public.status_task_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE public.status_task_id_seq;
       public       postgres    false    202            2           0    0    status_task_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE public.status_task_id_seq OWNED BY public.status_task.id;
            public       postgres    false    201            �            1259    16839    task    TABLE     k  CREATE TABLE public.task (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    deadline character varying(255) NOT NULL,
    user_id bigint NOT NULL,
    category_id bigint NOT NULL,
    status_id bigint NOT NULL,
    created_at character varying(255),
    updated_at character varying(255)
);
    DROP TABLE public.task;
       public         postgres    false            �            1259    16837    task_id_seq    SEQUENCE     �   CREATE SEQUENCE public.task_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.task_id_seq;
       public       postgres    false    200            3           0    0    task_id_seq    SEQUENCE OWNED BY     ;   ALTER SEQUENCE public.task_id_seq OWNED BY public.task.id;
            public       postgres    false    199            �            1259    16828    users    TABLE       CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    created_at character varying(255),
    updated_at character varying(255)
);
    DROP TABLE public.users;
       public         postgres    false            �            1259    16826    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public       postgres    false    198            4           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
            public       postgres    false    197            �
           2604    16864    category id    DEFAULT     j   ALTER TABLE ONLY public.category ALTER COLUMN id SET DEFAULT nextval('public.category_id_seq'::regclass);
 :   ALTER TABLE public.category ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    203    204    204            �
           2604    16853    status_task id    DEFAULT     p   ALTER TABLE ONLY public.status_task ALTER COLUMN id SET DEFAULT nextval('public.status_task_id_seq'::regclass);
 =   ALTER TABLE public.status_task ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    202    201    202            �
           2604    16842    task id    DEFAULT     b   ALTER TABLE ONLY public.task ALTER COLUMN id SET DEFAULT nextval('public.task_id_seq'::regclass);
 6   ALTER TABLE public.task ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    199    200    200            �
           2604    16831    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    198    197    198            *          0    16861    category 
   TABLE DATA               D   COPY public.category (id, name, created_at, updated_at) FROM stdin;
    public       postgres    false    204   )       "          0    16818    gorp_migrations 
   TABLE DATA               9   COPY public.gorp_migrations (id, applied_at) FROM stdin;
    public       postgres    false    196   P)       (          0    16850    status_task 
   TABLE DATA               I   COPY public.status_task (id, status, created_at, updated_at) FROM stdin;
    public       postgres    false    202   �)       &          0    16839    task 
   TABLE DATA               y   COPY public.task (id, title, description, deadline, user_id, category_id, status_id, created_at, updated_at) FROM stdin;
    public       postgres    false    200   �)       $          0    16828    users 
   TABLE DATA               R   COPY public.users (id, name, email, password, created_at, updated_at) FROM stdin;
    public       postgres    false    198   Y*       5           0    0    category_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.category_id_seq', 3, true);
            public       postgres    false    203            6           0    0    status_task_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.status_task_id_seq', 4, true);
            public       postgres    false    201            7           0    0    task_id_seq    SEQUENCE SET     9   SELECT pg_catalog.setval('public.task_id_seq', 5, true);
            public       postgres    false    199            8           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 2, true);
            public       postgres    false    197            �
           2606    16869    category category_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.category DROP CONSTRAINT category_pkey;
       public         postgres    false    204            �
           2606    16825 $   gorp_migrations gorp_migrations_pkey 
   CONSTRAINT     b   ALTER TABLE ONLY public.gorp_migrations
    ADD CONSTRAINT gorp_migrations_pkey PRIMARY KEY (id);
 N   ALTER TABLE ONLY public.gorp_migrations DROP CONSTRAINT gorp_migrations_pkey;
       public         postgres    false    196            �
           2606    16858    status_task status_task_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.status_task
    ADD CONSTRAINT status_task_pkey PRIMARY KEY (id);
 F   ALTER TABLE ONLY public.status_task DROP CONSTRAINT status_task_pkey;
       public         postgres    false    202            �
           2606    16847    task task_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.task DROP CONSTRAINT task_pkey;
       public         postgres    false    200            �
           2606    16836    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public         postgres    false    198            �
           2606    16880    task task_category_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.category(id);
 D   ALTER TABLE ONLY public.task DROP CONSTRAINT task_category_id_fkey;
       public       postgres    false    2725    200    204            �
           2606    16875    task task_status_id_fkey    FK CONSTRAINT        ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_status_id_fkey FOREIGN KEY (status_id) REFERENCES public.status_task(id);
 B   ALTER TABLE ONLY public.task DROP CONSTRAINT task_status_id_fkey;
       public       postgres    false    200    202    2723            �
           2606    16870    task task_user_id_fkey    FK CONSTRAINT     u   ALTER TABLE ONLY public.task
    ADD CONSTRAINT task_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);
 @   ALTER TABLE ONLY public.task DROP CONSTRAINT task_user_id_fkey;
       public       postgres    false    200    198    2719            *   3   x�3��IM,���K�4202�5"3C3+S+CK1s+S+c�=... y��      "   7   x�3����,�L,I�+.��4202�5"3CS+C#+#=S3sc#ms�=... W��      (   N   x�3�LI�I�U(H�KO-�JL��4202�5"3C3+SK+cs1+s+#s.c��Ԝ���LtY+C 1�=... D�Z      &   T   x�3�)MO,V0��M�KMO-�J�N��LJ-�/*��4�4BN.c�2C�҂�ĒTL�p�FFF��@d�`hae`dej����� ���      $   �   x�m�1��0@��W8��+T��k��3�TsK!X<�pr���k�d4y�[��(�ʬ2z�Ku8�iU"�*���y�q��7	M�i�bK���M���g<�	���V^��o������\?��wST�����vJ�я��qq�bэ��Hn�u�.t�.�Τ��po��X�1>���~9 ��~�H��;�l��A G�     