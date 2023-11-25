PGDMP      0             
    {            app    16.1    16.1 8    F           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            G           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            H           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            I           1262    16384    app    DATABASE     o   CREATE DATABASE app WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.UTF-8';
    DROP DATABASE app;
                app    false                        2615    2200    public    SCHEMA     2   -- *not* creating schema, since initdb creates it
 2   -- *not* dropping schema, since initdb creates it
                app    false            �            1255    16453    table_notify()    FUNCTION     �  CREATE FUNCTION public.table_notify() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE
  channel TEXT;
  old_row JSON;
  new_row JSON;
  notification JSON;
  xmin BIGINT;
  _indices TEXT [];
  _primary_keys TEXT [];
  _foreign_keys TEXT [];

BEGIN
    -- database is also the channel name.
    channel := CURRENT_DATABASE();

    IF TG_OP = 'DELETE' THEN

        SELECT primary_keys, indices
        INTO _primary_keys, _indices
        FROM public._view
        WHERE table_name = TG_TABLE_NAME;

        old_row = ROW_TO_JSON(OLD);
        old_row := (
            SELECT JSONB_OBJECT_AGG(key, value)
            FROM JSON_EACH(old_row)
            WHERE key = ANY(_primary_keys)
        );
        xmin := OLD.xmin;
    ELSE
        IF TG_OP <> 'TRUNCATE' THEN

            SELECT primary_keys, foreign_keys, indices
            INTO _primary_keys, _foreign_keys, _indices
            FROM public._view
            WHERE table_name = TG_TABLE_NAME;

            new_row = ROW_TO_JSON(NEW);
            new_row := (
                SELECT JSONB_OBJECT_AGG(key, value)
                FROM JSON_EACH(new_row)
                WHERE key = ANY(_primary_keys || _foreign_keys)
            );
            IF TG_OP = 'UPDATE' THEN
                old_row = ROW_TO_JSON(OLD);
                old_row := (
                    SELECT JSONB_OBJECT_AGG(key, value)
                    FROM JSON_EACH(old_row)
                    WHERE key = ANY(_primary_keys || _foreign_keys)
                );
            END IF;
            xmin := NEW.xmin;
        END IF;
    END IF;

    -- construct the notification as a JSON object.
    notification = JSON_BUILD_OBJECT(
        'xmin', xmin,
        'new', new_row,
        'old', old_row,
        'indices', _indices,
        'tg_op', TG_OP,
        'table', TG_TABLE_NAME,
        'schema', TG_TABLE_SCHEMA
    );

    -- Notify/Listen updates occur asynchronously,
    -- so this doesn't block the Postgres trigger procedure.
    PERFORM PG_NOTIFY(channel, notification::TEXT);

  RETURN NEW;
END;
$$;
 %   DROP FUNCTION public.table_notify();
       public          postgres    false    5            �            1259    16454    _view    MATERIALIZED VIEW     �  CREATE MATERIALIZED VIEW public._view AS
 SELECT table_name,
    primary_keys,
    foreign_keys,
    indices
   FROM ( VALUES ('hashtags'::text,ARRAY['id'::text],NULL::text[],ARRAY['projects'::text]), ('project_hashtags'::text,ARRAY['hashtag_id'::text, 'project_id'::text],ARRAY['hashtag_id'::text, 'project_id'::text],ARRAY['projects'::text]), ('projects'::text,ARRAY['id'::text],NULL::text[],ARRAY['projects'::text]), ('user_projects'::text,ARRAY['user_id'::text, 'project_id'::text],ARRAY['user_id'::text, 'project_id'::text],ARRAY['projects'::text]), ('users'::text,ARRAY['id'::text],NULL::text[],ARRAY['projects'::text])) t(table_name, primary_keys, foreign_keys, indices)
  WITH NO DATA;
 %   DROP MATERIALIZED VIEW public._view;
       public         heap    app    false    5            �            1259    16386    hashtags    TABLE     �   CREATE TABLE public.hashtags (
    id bigint NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone NOT NULL
);
    DROP TABLE public.hashtags;
       public         heap    app    false    5            �            1259    16391    hashtags_id_seq    SEQUENCE     x   CREATE SEQUENCE public.hashtags_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.hashtags_id_seq;
       public          app    false    215    5            J           0    0    hashtags_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.hashtags_id_seq OWNED BY public.hashtags.id;
          public          app    false    216            �            1259    16392    project_hashtags    TABLE     i   CREATE TABLE public.project_hashtags (
    hashtag_id bigint NOT NULL,
    project_id bigint NOT NULL
);
 $   DROP TABLE public.project_hashtags;
       public         heap    app    false    5            �            1259    16395    projects    TABLE     �   CREATE TABLE public.projects (
    id bigint NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    slug text DEFAULT ''::text NOT NULL,
    description text DEFAULT ''::text NOT NULL,
    created_at timestamp with time zone NOT NULL
);
    DROP TABLE public.projects;
       public         heap    app    false    5            �            1259    16403    projects_id_seq    SEQUENCE     x   CREATE SEQUENCE public.projects_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.projects_id_seq;
       public          app    false    218    5            K           0    0    projects_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.projects_id_seq OWNED BY public.projects.id;
          public          app    false    219            �            1259    16404    user_projects    TABLE     c   CREATE TABLE public.user_projects (
    project_id bigint NOT NULL,
    user_id bigint NOT NULL
);
 !   DROP TABLE public.user_projects;
       public         heap    app    false    5            �            1259    16407    users    TABLE     �   CREATE TABLE public.users (
    id bigint NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    created_at timestamp with time zone NOT NULL
);
    DROP TABLE public.users;
       public         heap    app    false    5            �            1259    16413    users_id_seq    SEQUENCE     u   CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          app    false    221    5            L           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          app    false    222            �           2604    16414    hashtags id    DEFAULT     j   ALTER TABLE ONLY public.hashtags ALTER COLUMN id SET DEFAULT nextval('public.hashtags_id_seq'::regclass);
 :   ALTER TABLE public.hashtags ALTER COLUMN id DROP DEFAULT;
       public          app    false    216    215            �           2604    16415    projects id    DEFAULT     j   ALTER TABLE ONLY public.projects ALTER COLUMN id SET DEFAULT nextval('public.projects_id_seq'::regclass);
 :   ALTER TABLE public.projects ALTER COLUMN id DROP DEFAULT;
       public          app    false    219    218            �           2604    16416    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          app    false    222    221            ;          0    16386    hashtags 
   TABLE DATA           8   COPY public.hashtags (id, name, created_at) FROM stdin;
    public          app    false    215   �H       =          0    16392    project_hashtags 
   TABLE DATA           B   COPY public.project_hashtags (hashtag_id, project_id) FROM stdin;
    public          app    false    217   WK       >          0    16395    projects 
   TABLE DATA           K   COPY public.projects (id, name, slug, description, created_at) FROM stdin;
    public          app    false    218   pL       @          0    16404    user_projects 
   TABLE DATA           <   COPY public.user_projects (project_id, user_id) FROM stdin;
    public          app    false    220   JR       A          0    16407    users 
   TABLE DATA           5   COPY public.users (id, name, created_at) FROM stdin;
    public          app    false    221   cS       M           0    0    hashtags_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.hashtags_id_seq', 100, true);
          public          app    false    216            N           0    0    projects_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.projects_id_seq', 50, true);
          public          app    false    219            O           0    0    users_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.users_id_seq', 100, true);
          public          app    false    222            �           2606    16418    hashtags hashtags_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.hashtags
    ADD CONSTRAINT hashtags_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.hashtags DROP CONSTRAINT hashtags_pkey;
       public            app    false    215            �           2606    16420 &   project_hashtags project_hashtags_pkey 
   CONSTRAINT     x   ALTER TABLE ONLY public.project_hashtags
    ADD CONSTRAINT project_hashtags_pkey PRIMARY KEY (hashtag_id, project_id);
 P   ALTER TABLE ONLY public.project_hashtags DROP CONSTRAINT project_hashtags_pkey;
       public            app    false    217    217            �           2606    16422    projects projects_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.projects DROP CONSTRAINT projects_pkey;
       public            app    false    218            �           2606    16424     user_projects user_projects_pkey 
   CONSTRAINT     o   ALTER TABLE ONLY public.user_projects
    ADD CONSTRAINT user_projects_pkey PRIMARY KEY (project_id, user_id);
 J   ALTER TABLE ONLY public.user_projects DROP CONSTRAINT user_projects_pkey;
       public            app    false    220    220            �           2606    16426    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            app    false    221            �           1259    16460    _idx    INDEX     C   CREATE UNIQUE INDEX _idx ON public._view USING btree (table_name);
    DROP INDEX public._idx;
       public            app    false    223            �           1259    16427    hashtags_name_idx    INDEX     M   CREATE UNIQUE INDEX hashtags_name_idx ON public.hashtags USING btree (name);
 %   DROP INDEX public.hashtags_name_idx;
       public            app    false    215            �           1259    16428    project_hashtags_project_id_idx    INDEX     b   CREATE INDEX project_hashtags_project_id_idx ON public.project_hashtags USING btree (project_id);
 3   DROP INDEX public.project_hashtags_project_id_idx;
       public            app    false    217            �           1259    16429    projects_name_idx    INDEX     F   CREATE INDEX projects_name_idx ON public.projects USING btree (name);
 %   DROP INDEX public.projects_name_idx;
       public            app    false    218            �           1259    16430    projects_slug_idx    INDEX     F   CREATE INDEX projects_slug_idx ON public.projects USING btree (slug);
 %   DROP INDEX public.projects_slug_idx;
       public            app    false    218            �           1259    16431    user_projects_user_id_idx    INDEX     V   CREATE INDEX user_projects_user_id_idx ON public.user_projects USING btree (user_id);
 -   DROP INDEX public.user_projects_user_id_idx;
       public            app    false    220            �           1259    16432    users_name_idx    INDEX     G   CREATE UNIQUE INDEX users_name_idx ON public.users USING btree (name);
 "   DROP INDEX public.users_name_idx;
       public            app    false    221            �           2620    16461    hashtags hashtags_notify    TRIGGER     �   CREATE TRIGGER hashtags_notify AFTER INSERT OR DELETE OR UPDATE ON public.hashtags FOR EACH ROW EXECUTE FUNCTION public.table_notify();
 1   DROP TRIGGER hashtags_notify ON public.hashtags;
       public          app    false    215    224            �           2620    16462    hashtags hashtags_truncate    TRIGGER     ~   CREATE TRIGGER hashtags_truncate AFTER TRUNCATE ON public.hashtags FOR EACH STATEMENT EXECUTE FUNCTION public.table_notify();
 3   DROP TRIGGER hashtags_truncate ON public.hashtags;
       public          app    false    224    215            �           2620    16463 (   project_hashtags project_hashtags_notify    TRIGGER     �   CREATE TRIGGER project_hashtags_notify AFTER INSERT OR DELETE OR UPDATE ON public.project_hashtags FOR EACH ROW EXECUTE FUNCTION public.table_notify();
 A   DROP TRIGGER project_hashtags_notify ON public.project_hashtags;
       public          app    false    224    217            �           2620    16464 *   project_hashtags project_hashtags_truncate    TRIGGER     �   CREATE TRIGGER project_hashtags_truncate AFTER TRUNCATE ON public.project_hashtags FOR EACH STATEMENT EXECUTE FUNCTION public.table_notify();
 C   DROP TRIGGER project_hashtags_truncate ON public.project_hashtags;
       public          app    false    224    217            �           2620    16465    projects projects_notify    TRIGGER     �   CREATE TRIGGER projects_notify AFTER INSERT OR DELETE OR UPDATE ON public.projects FOR EACH ROW EXECUTE FUNCTION public.table_notify();
 1   DROP TRIGGER projects_notify ON public.projects;
       public          app    false    218    224            �           2620    16466    projects projects_truncate    TRIGGER     ~   CREATE TRIGGER projects_truncate AFTER TRUNCATE ON public.projects FOR EACH STATEMENT EXECUTE FUNCTION public.table_notify();
 3   DROP TRIGGER projects_truncate ON public.projects;
       public          app    false    224    218            �           2620    16467 "   user_projects user_projects_notify    TRIGGER     �   CREATE TRIGGER user_projects_notify AFTER INSERT OR DELETE OR UPDATE ON public.user_projects FOR EACH ROW EXECUTE FUNCTION public.table_notify();
 ;   DROP TRIGGER user_projects_notify ON public.user_projects;
       public          app    false    220    224            �           2620    16468 $   user_projects user_projects_truncate    TRIGGER     �   CREATE TRIGGER user_projects_truncate AFTER TRUNCATE ON public.user_projects FOR EACH STATEMENT EXECUTE FUNCTION public.table_notify();
 =   DROP TRIGGER user_projects_truncate ON public.user_projects;
       public          app    false    224    220            �           2620    16469    users users_notify    TRIGGER     �   CREATE TRIGGER users_notify AFTER INSERT OR DELETE OR UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.table_notify();
 +   DROP TRIGGER users_notify ON public.users;
       public          app    false    221    224            �           2620    16470    users users_truncate    TRIGGER     x   CREATE TRIGGER users_truncate AFTER TRUNCATE ON public.users FOR EACH STATEMENT EXECUTE FUNCTION public.table_notify();
 -   DROP TRIGGER users_truncate ON public.users;
       public          app    false    221    224            �           2606    16433 1   project_hashtags project_hashtags_hashtag_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.project_hashtags
    ADD CONSTRAINT project_hashtags_hashtag_id_fkey FOREIGN KEY (hashtag_id) REFERENCES public.hashtags(id);
 [   ALTER TABLE ONLY public.project_hashtags DROP CONSTRAINT project_hashtags_hashtag_id_fkey;
       public          app    false    217    3214    215            �           2606    16438 1   project_hashtags project_hashtags_project_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.project_hashtags
    ADD CONSTRAINT project_hashtags_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id);
 [   ALTER TABLE ONLY public.project_hashtags DROP CONSTRAINT project_hashtags_project_id_fkey;
       public          app    false    3220    217    218            �           2606    16443 +   user_projects user_projects_project_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.user_projects
    ADD CONSTRAINT user_projects_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id);
 U   ALTER TABLE ONLY public.user_projects DROP CONSTRAINT user_projects_project_id_fkey;
       public          app    false    3220    220    218            �           2606    16448 (   user_projects user_projects_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.user_projects
    ADD CONSTRAINT user_projects_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);
 R   ALTER TABLE ONLY public.user_projects DROP CONSTRAINT user_projects_user_id_fkey;
       public          app    false    220    221    3227            C           0    16454    _view    MATERIALIZED VIEW DATA     (   REFRESH MATERIALIZED VIEW public._view;
          public          app    false    223    3397            ;   �  x��VA�� <+����L0�-{!6q���<���*y@s�B�[ݒ �������C�_�}��k4�V���p���m����b�zG.�4�y?|�5'����`G�� �%�Z�s����k�6���-�%(���/$���Rn�i1�|�DL�ЬȺ��#@Lt��4T_���Y~g�G%<}��If;�����?`�Rr%������_J��#�x� DS.K(3�����ӑ����3�����v�@����s�x��`�=��b�}L�e�4J�[�f�Įk�&�75�#���2�q"��:W�y�Ɔ����qV��(#���%���;�(�{��Q}��9�g_n�_)���8эI�C3����,�{���z�:���8�
5���	�/Ig��(�J��j���`�jMK�[>���=ޘ�2h̓~㒻opf�eV���rt�)�]6t����A��[���L�P�y_��ޛ�=i
�:�&�\��7�1�>(�y��+'Β�R�ơ�,g� #�I�YJg��b�w8"f�2=���1� M�8q��K��R�['��n��v���VpԳ��=d�[F���;v6�_��W��莝�_s��Ѱ�z��Z�(n���~���b�p�����t�Oy߂p������pL�����M�=�����g���Vs�Rg�a�%�����>/��?      =   	  x��[�dQB�o�c�x��u�)����&��b�9.���=أ�hsZ������hk�����������v�!�<�	Oz����|�_�����ac8�{��F"�Q�b4���"�q�c>�cҘ�tf ��/2��,f#�9�a.r��<�C=���������㩟O���'T?�Q��E�"j��ZD-�Q�hD4"��FD#�шhD4"Z��VD+�ъhE�"Z��ND'�щ�Dt":��ND'"{2��G�?pTN      >   �  x��Y�n9}6_a�Z��}㍄KXD��Vڗj�:1��-۝���[� = ��@��!RO�cW�:�.�=󈖿s>:��:Cݱw���m��������|������?ô�p�@o�YQ>��E���a%�f�-�?��^�����2�Y��E�=���?�稸u�;kf>�G.���k��n�>F�Gn��e�+�V�l�F�6d��݆2bq)�˳���`���2�v_!������u�~���I����ri�k,�9��M��k��g~�%z��f��cB�E-Zv�`W΢�8�F�SmO{y�`�ι���q�����O��J������L������q�a�gÎ�΢g� �`�F~���э��{4�]�[ z�� �U�gm���C$g�+.O'{�5��Ń��<��,p���;5I��'��i�#���D�Q�񷻢������9����}����K ���'ڂI�lv�ѱ�o���G���OZh+1�G�,�n�p�X��#���%K#m�bAV!D>z�Pk��%�j��3D�c~�^Niq'd'�/$D���p5m6�¤�d��	��ӄr�G��ڪ)D?�}�TR|�P���%���['�9��y�˙�3��H���M�W͒k�	�Ƒe�X�Ex)��@B�x�)=r��6�:�m^,HR���B��p�ǋS��8w�@A����n�{<1~t�z�Ҋ��]���H�8��P�H�~C�Z�!�&f�m��~'5{m���� D�do�<M(�C9y�3s��o4Iiϼ;G���,$�]	+ŞhK�RͩU5�\�ꚓo垷<�ٗ"���(֤�����Ϛ��˟��T
�",�y}uE�I6�l��)V�_�9�YV�H��yƗ{�����FUR�͝dNY�sn��E�^�t�4�����ײ��݌ટ}�Y�g
k�3ǟjk�q!x�l�}-)-_g]��Ɇ�C��������-�.��u�|��J�����cf�������eE�{��j�;�q�ӶQ�α<���T�X�q��Ê�U�j���Ȫ��)��d�aV~��\���_>L�������$cr�p�^�t�T�Y���8���O~$ݼ��Z˰�t�rw��^��{�4�(��n�N�l��i� ���0�U]����!Ƨ����?�	�FA�K��\c~lPS�	�~P��|��;Z�����:&��*n�yk)ۧ��Y��<����LT[3��36�X�]����eo�I����`Mݵ��_����m)DPγ��d�T?s|�ؾ!fQ�����Av���;0�����d�eM=��7��ʭ����VL]>|�LD@yG��g��<���H���^�凬��g���h���l#\���J,��-n,bm�k�$DV�(j���٦lٻ����<�cU���|kQ��摡��*�'�].J*n��4�)X9���L!P����*K�i�cgg&�Rdշf���ݻ�/%|߭      @   	  x��[�dQB�o�c�x��u�)����&��b�9.���=أ�hsZ������hk�����������v�!�<�	Oz����|�_�����ac8�{��F"�Q�b4���"�q�c>�cҘ�tf ��/2��,f#�9�a.r��<�C}��M�l�S����ϧ~@�*��ZD-�Q��E�"j��ZD#�шhD4"��FD#�ъhE�"Z��VD+�ъhEt":��ND'�щ�Dt":�{��G�? vTN      A   �  x��W�r�8<�_������&y�-9�פ��L"$F ���m�����I�.|��F,��-�,r߼x�*�DI�)�?%���/y�%+?WI�g�%�=��'�SqK��'����Z&g�Ƽ�W�9l.�޹�Y�<8c�9t!�������V-��zy�͋b?�O�m�#�q�Z�u7+#����8h�[��"o��׌c��1���+g&b��"��9��)NPF� ���<q&�h1� o�9��"�odùcOư�B�̝��FV0q)6//4ʽ�޽��Jl��J(�+O,���W�5@$�=5Z�g2G�Yl,��W}V�>I�F*�[��C��B��!?��I2q9�+�����bG3��@�u�Գ�B�i���m&�'A��Җ?hp<	Iyi� ��]!�_�no02v
i$�=��iHc���8����Ґ�����&�:,M�y�̈	��1^3)(k��q~�X��f�ʝ��H�
�1T(�5,0�5a e�ȇaZ�sgø%��qP���؜�6�®��E�N?�q(;�,��-/f���G�i;�#��R$a��;���d⠛�"/h�z���,�w�;��Bl~�$�����q�$�smh�/\�K�%�Sή�MV�����!�S�H\��d�����bc(��������^�b?�Y:�Xh*6�_LW�^y&~`���Ƭ� ��-A]ty!��Q[�H�~k��^�l���WP�#y��ޙ������΃��,E$����,��!hB��+!V��I���[)_E
�Y�h돯�E-� g��K���6�8ivdE�b��	�ݍ�����=U�l��
�L�<OI5|q�,�+�u`�K�5��n-�������;�2ƺPV�q�n�2wa	����\��o����)| �`&*ك;����9�'�A	ߵK�d(����߂e��0NK�������"�V�Κ�J�5�)���4��*B��B�k�*�&���^O+�� ���V�i��\����~�L�M����v�>�#�_�����'OU���"�h��W�x�4H�@��fs��� �>�6x�T�xZ���U�]��c+^15�:D��5�g�<~����T<�e��n՚��,��	%@����5uݺy�$��5�2�`3�S���e�c��k;~hUxv�i�ִh��yj��Ew�	����*�����2Â��|vv�6�	�     