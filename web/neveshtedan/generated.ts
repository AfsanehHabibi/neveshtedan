import { GraphQLResolveInfo } from 'graphql';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
export type RequireFields<T, K extends keyof T> = Omit<T, K> & { [P in K]-?: NonNullable<T[P]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type Login = {
  password: Scalars['String']['input'];
  username: Scalars['String']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createUser: Scalars['String']['output'];
  createWritingEntry: WritingEntry;
  login: Scalars['String']['output'];
  refreshToken: Scalars['String']['output'];
};


export type MutationCreateUserArgs = {
  input: NewUser;
};


export type MutationCreateWritingEntryArgs = {
  input: NewWritingEntry;
};


export type MutationLoginArgs = {
  input: Login;
};


export type MutationRefreshTokenArgs = {
  input: RefreshTokenInput;
};

export type NewUser = {
  password: Scalars['String']['input'];
  username: Scalars['String']['input'];
};

export type NewWritingEntry = {
  fields: Array<NewWritingEntryField>;
  templateId: Scalars['Int']['input'];
  userId: Scalars['Int']['input'];
};

export type NewWritingEntryField = {
  name: Scalars['String']['input'];
  value?: InputMaybe<Scalars['String']['input']>;
};

export type Query = {
  __typename?: 'Query';
  entries: Array<WritingEntry>;
  templates: Array<WritingTemplate>;
  writingTemplate?: Maybe<WritingTemplate>;
};


export type QueryWritingTemplateArgs = {
  id: Scalars['Int']['input'];
};

export type RefreshTokenInput = {
  token: Scalars['String']['input'];
};

export type User = {
  __typename?: 'User';
  id: Scalars['Int']['output'];
  username: Scalars['String']['output'];
};

export type WritingEntry = {
  __typename?: 'WritingEntry';
  fields: Array<WritingEntryField>;
  id: Scalars['Int']['output'];
  templateId: Scalars['Int']['output'];
  userId: Scalars['Int']['output'];
};

export type WritingEntryField = {
  __typename?: 'WritingEntryField';
  name: Scalars['String']['output'];
  value?: Maybe<Scalars['String']['output']>;
};

export type WritingTemplate = {
  __typename?: 'WritingTemplate';
  fields: Array<Scalars['String']['output']>;
  id: Scalars['Int']['output'];
  title: Scalars['String']['output'];
};



export type ResolverTypeWrapper<T> = Promise<T> | T;


export type ResolverWithResolve<TResult, TParent, TContext, TArgs> = {
  resolve: ResolverFn<TResult, TParent, TContext, TArgs>;
};
export type Resolver<TResult, TParent = {}, TContext = {}, TArgs = {}> = ResolverFn<TResult, TParent, TContext, TArgs> | ResolverWithResolve<TResult, TParent, TContext, TArgs>;

export type ResolverFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => Promise<TResult> | TResult;

export type SubscriptionSubscribeFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => AsyncIterable<TResult> | Promise<AsyncIterable<TResult>>;

export type SubscriptionResolveFn<TResult, TParent, TContext, TArgs> = (
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => TResult | Promise<TResult>;

export interface SubscriptionSubscriberObject<TResult, TKey extends string, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<{ [key in TKey]: TResult }, TParent, TContext, TArgs>;
  resolve?: SubscriptionResolveFn<TResult, { [key in TKey]: TResult }, TContext, TArgs>;
}

export interface SubscriptionResolverObject<TResult, TParent, TContext, TArgs> {
  subscribe: SubscriptionSubscribeFn<any, TParent, TContext, TArgs>;
  resolve: SubscriptionResolveFn<TResult, any, TContext, TArgs>;
}

export type SubscriptionObject<TResult, TKey extends string, TParent, TContext, TArgs> =
  | SubscriptionSubscriberObject<TResult, TKey, TParent, TContext, TArgs>
  | SubscriptionResolverObject<TResult, TParent, TContext, TArgs>;

export type SubscriptionResolver<TResult, TKey extends string, TParent = {}, TContext = {}, TArgs = {}> =
  | ((...args: any[]) => SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>)
  | SubscriptionObject<TResult, TKey, TParent, TContext, TArgs>;

export type TypeResolveFn<TTypes, TParent = {}, TContext = {}> = (
  parent: TParent,
  context: TContext,
  info: GraphQLResolveInfo
) => Maybe<TTypes> | Promise<Maybe<TTypes>>;

export type IsTypeOfResolverFn<T = {}, TContext = {}> = (obj: T, context: TContext, info: GraphQLResolveInfo) => boolean | Promise<boolean>;

export type NextResolverFn<T> = () => Promise<T>;

export type DirectiveResolverFn<TResult = {}, TParent = {}, TContext = {}, TArgs = {}> = (
  next: NextResolverFn<TResult>,
  parent: TParent,
  args: TArgs,
  context: TContext,
  info: GraphQLResolveInfo
) => TResult | Promise<TResult>;



/** Mapping between all available schema types and the resolvers types */
export type ResolversTypes = {
  Boolean: ResolverTypeWrapper<Scalars['Boolean']['output']>;
  Int: ResolverTypeWrapper<Scalars['Int']['output']>;
  Login: Login;
  Mutation: ResolverTypeWrapper<{}>;
  NewUser: NewUser;
  NewWritingEntry: NewWritingEntry;
  NewWritingEntryField: NewWritingEntryField;
  Query: ResolverTypeWrapper<{}>;
  RefreshTokenInput: RefreshTokenInput;
  String: ResolverTypeWrapper<Scalars['String']['output']>;
  User: ResolverTypeWrapper<User>;
  WritingEntry: ResolverTypeWrapper<WritingEntry>;
  WritingEntryField: ResolverTypeWrapper<WritingEntryField>;
  WritingTemplate: ResolverTypeWrapper<WritingTemplate>;
};

/** Mapping between all available schema types and the resolvers parents */
export type ResolversParentTypes = {
  Boolean: Scalars['Boolean']['output'];
  Int: Scalars['Int']['output'];
  Login: Login;
  Mutation: {};
  NewUser: NewUser;
  NewWritingEntry: NewWritingEntry;
  NewWritingEntryField: NewWritingEntryField;
  Query: {};
  RefreshTokenInput: RefreshTokenInput;
  String: Scalars['String']['output'];
  User: User;
  WritingEntry: WritingEntry;
  WritingEntryField: WritingEntryField;
  WritingTemplate: WritingTemplate;
};

export type MutationResolvers<ContextType = any, ParentType extends ResolversParentTypes['Mutation'] = ResolversParentTypes['Mutation']> = {
  createUser?: Resolver<ResolversTypes['String'], ParentType, ContextType, RequireFields<MutationCreateUserArgs, 'input'>>;
  createWritingEntry?: Resolver<ResolversTypes['WritingEntry'], ParentType, ContextType, RequireFields<MutationCreateWritingEntryArgs, 'input'>>;
  login?: Resolver<ResolversTypes['String'], ParentType, ContextType, RequireFields<MutationLoginArgs, 'input'>>;
  refreshToken?: Resolver<ResolversTypes['String'], ParentType, ContextType, RequireFields<MutationRefreshTokenArgs, 'input'>>;
};

export type QueryResolvers<ContextType = any, ParentType extends ResolversParentTypes['Query'] = ResolversParentTypes['Query']> = {
  entries?: Resolver<Array<ResolversTypes['WritingEntry']>, ParentType, ContextType>;
  templates?: Resolver<Array<ResolversTypes['WritingTemplate']>, ParentType, ContextType>;
  writingTemplate?: Resolver<Maybe<ResolversTypes['WritingTemplate']>, ParentType, ContextType, RequireFields<QueryWritingTemplateArgs, 'id'>>;
};

export type UserResolvers<ContextType = any, ParentType extends ResolversParentTypes['User'] = ResolversParentTypes['User']> = {
  id?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  username?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type WritingEntryResolvers<ContextType = any, ParentType extends ResolversParentTypes['WritingEntry'] = ResolversParentTypes['WritingEntry']> = {
  fields?: Resolver<Array<ResolversTypes['WritingEntryField']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  templateId?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  userId?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type WritingEntryFieldResolvers<ContextType = any, ParentType extends ResolversParentTypes['WritingEntryField'] = ResolversParentTypes['WritingEntryField']> = {
  name?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  value?: Resolver<Maybe<ResolversTypes['String']>, ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type WritingTemplateResolvers<ContextType = any, ParentType extends ResolversParentTypes['WritingTemplate'] = ResolversParentTypes['WritingTemplate']> = {
  fields?: Resolver<Array<ResolversTypes['String']>, ParentType, ContextType>;
  id?: Resolver<ResolversTypes['Int'], ParentType, ContextType>;
  title?: Resolver<ResolversTypes['String'], ParentType, ContextType>;
  __isTypeOf?: IsTypeOfResolverFn<ParentType, ContextType>;
};

export type Resolvers<ContextType = any> = {
  Mutation?: MutationResolvers<ContextType>;
  Query?: QueryResolvers<ContextType>;
  User?: UserResolvers<ContextType>;
  WritingEntry?: WritingEntryResolvers<ContextType>;
  WritingEntryField?: WritingEntryFieldResolvers<ContextType>;
  WritingTemplate?: WritingTemplateResolvers<ContextType>;
};

